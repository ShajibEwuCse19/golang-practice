package cmd

import (
	"ecom-project-improved/config"
	"ecom-project-improved/infra/db"
	"ecom-project-improved/repo"
	"ecom-project-improved/rest"
	srvProduct "ecom-project-improved/services/product"
	srvUser "ecom-project-improved/services/user"
	productHandler "ecom-project-improved/rest/handlers/product"
	userHandler "ecom-project-improved/rest/handlers/user"
	"ecom-project-improved/utils"
	"fmt"
	"os"
	"strconv"
)

func Serve() {
	config := config.GetConfig()
	fmt.Println("Version:", config.Version)
	fmt.Println("Service Name:", config.ServiceName)
	fmt.Println("Http Port:", config.HttpPort)
	utils.FreePort(":" + strconv.Itoa(config.HttpPort)) // Automatically check and kill any old or suspended process holding this port

	dbCon, err := db.NewConnection(config) // connect with postgres db using sqlx. Now, dbCon.DB is our database
	if err != nil {
		fmt.Println("db connection err:", err)
		os.Exit(1)
	}

	// Migrate Database
	err = db.MigrateDB(dbCon, "./migrations") //For migration, we need live postgres db [dbCon.DB], and migration query directory path.
	if err != nil {
		fmt.Println("db migration err:", err)
		os.Exit(1)
	}

	// 1. REPOSITORIES:
	// - Why needed: To execute actual SQL queries (INSERT, SELECT, UPDATE, DELETE).
	// - Why pass dbCon: The repo needs a live PostgreSQL database connection handle (*sqlx.DB) to talk to the DB.
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	// 2. DOMAINS / SERVICES:
	// - Why needed: To hold core business logic, rules, and validations (e.g., password hashing, business checks).
	// - Why pass userRepo or productRepo: The service layer manages business rules but does NOT write raw SQL.
	//   It delegates data saving/fetching to the repo interface.
	userSvc := srvUser.NewService(userRepo)
	productSvc := srvProduct.NewService(productRepo)

	// 3. HANDLERS (HTTP Controllers):
	// - Why needed: To handle the HTTP transport layer (read HTTP requests, decode JSON from r.Body, and write HTTP status/JSON responses).
	// - What exact thing uses userSvc or productSvc: Controller functions inside rest/handlers/... (e.g., h.svc.Create(...) in create_products.go,
	//   h.svc.Login(...) in login.go).
	// - Why pass userSvc or productSvc: Handlers must 'NOT talk to databases or perform business logic directly'.
	//   Once a valid HTTP request is parsed, the handler calls the service to execute the actual business operation.
	userHandler := userHandler.NewHandler(userSvc)
	productHandler := productHandler.NewHandler(productSvc)

	// 4. SERVER:
	// - Why needed: To setup the HTTP router (http.NewServeMux), wrap global middlewares (CORS, Logger), and listen on the TCP port (http.ListenAndServe).
	// - What exact thing uses productHandler and userHandler: Inside rest/server.go -> server.StartServer():
	//   - server.productHandler.RegisterRoutes(mux, manager) -> registers GET /products, POST /create-product, etc.
	//   - server.userHandler.RegisterRoutes(mux, manager) -> registers POST /register, POST /login, etc.
	// - Why pass handlers: The server does not know what API endpoints exist in your app. Passing handlers allows the server
	//   to mount their routes onto the mux router. Without handlers, the server is an empty shell returning 404 for all requests.
	server := rest.NewServer(productHandler, userHandler, config)
	server.StartServer()
}

/*

============================================================================================
                              ভিজ্যুয়াল ফ্লো ডায়াগ্রাম (Visual Flow Diagram)
============================================================================================

 [ Client (Postman / Browser / App) ]
                 │
                 │ 1. HTTP Request পাঠায় (যেমন: POST /create-product সাথে JSON Body)
                 ▼
     [ Server (rest/server.go) ]
                 │  - TCP পোর্টে কানেকশন ধরে এবং http.ServeMux রাউটারে পাঠায়
                 │ 2. রুট ম্যাচ করে মিডলওয়্যার পাইপলাইনে দেয়
                 ▼
 [ Middleware Layer (rest/middleware/) ]
                 │  - Cors, Logger, Preflight চেক করে
                 │  - JwtAuthMiddleware দিয়ে Bearer Token ভ্যালিডেট করে (ভুল হলে এখানেই 401 রিটার্ন)
                 │ 3. অথেনটিকেশন সফল হলে রিকোয়েস্ট হ্যান্ডলারে পৌঁছায়
                 ▼
   [ Handler (rest/handlers/product/create_products.go) ]
                 │  - r.Body থেকে JSON ডেটা ডিকোড করে Struct-এ নেয়
                 │  - DTO থেকে domain.Product মডেলে রূপান্তর করে
                 │ 4. বিজনেস অপারেশনের জন্য h.svc.Create(domain.Product) কল করে
                 ▼
   [ Service / Domain (services/product/service.go) ]
                 │  - কোর বিজনেস রুলস ও লজিক ভ্যালিডেট করে
                 │  - নিজে কোনো SQL রান না করে ডেটাবেজ সেভের দায়িত্ব দেয় Repo-কে
                 │ 5. ডাটাবেজ কন্ট্রাক্ট অনুযায়ী svc.productRepo.Create(p) কল করে
                 ▼
   [ Repository (repo/product.go) ]
                 │  - Raw SQL কুয়েরি প্রস্তুত করে (INSERT INTO products ... RETURNING id)
                 │  - dbCon.DB (*sqlx.DB) ব্যবহার করে Live PostgreSQL সার্ভারে কুয়েরি চালায়
                 │ 6. ডাটাবেজে SQL স্টেটমেন্ট এক্সিকিউট করে
                 ▼
        [( PostgreSQL Database )]
                 │  - টেবিলে নতুন সারি (Row) সেভ করে এবং নতুন অটো-ইনক্রিমেন্ট ID জেনারেট করে
                 │
                 │ 7. ডাটাবেজ থেকে নতুন ID ও রেকর্ড রেসপন্স হিসেবে ফিরে আসে
                 ▼
   [ Repository (repo/product.go) ]
                 │  - ডাটাবেজের রেকর্ডটি ডোমেন মডেলে স্ক্যান (&p.ID) করে
                 │ 8. পূর্ণাঙ্গ Product অবজেক্টটি সার্ভিসে রিটার্ন করে
                 ▼
   [ Service / Domain (services/product/service.go) ]
                 │  - কোনো অতিরিক্ত প্রসেসিং থাকলে তা সম্পন্ন করে
                 │ 9. প্রসেস করা ডেটা হ্যান্ডলারে রিটার্ন করে
                 ▼
   [ Handler (rest/handlers/product/create_products.go) ]
                 │  - utils.SendData() কল করে রেসপন্স হেডার (application/json) ও Status 201 সেট করে
                 │ 10. ক্লায়েন্টকে ফাইনাল JSON রেসপন্স রিটার্ন করে
                 ▼
 [ Client (Postman / Browser / App) ]
   (ক্লায়েন্ট নতুন ID সহ সফল 201 Created রেসপন্স পায় এবং ফ্লো সমাপ্ত হয়)


============================================================================================
                               ধাপে ধাপে বিস্তারিত ব্যাখ্যা
============================================================================================

ধাপ ১: ক্লায়েন্ট রিকোয়েস্ট (HTTP Request)
- ক্লায়েন্ট (যেমন: Postman, Frontend) সার্ভারের পোর্টে একটি HTTP রিকোয়েস্ট পাঠায় (যেমন: POST /create-product সাথে হেডার Authorization: Bearer <token> এবং JSON Body)।

ধাপ ২: সার্ভার ও রাউটিং (rest/server.go)
- http.ListenAndServe() কানেকশনটি গ্রহণ করে।
- server.StartServer() এ তৈরি হওয়া http.NewServeMux() রাউটার চেক করে দেখে কোন হ্যান্ডলার এই URL পাথের সাথে রেজিস্টার করা আছে।

ধাপ ৩: মিডলওয়্যার লেয়ার (rest/middleware/)
- গ্লোবাল মিডলওয়্যার (Logger, Cors, Preflight) রিকোয়েস্টটি প্রসেস করে।
- এরপর রুটের সাথে থাকা JwtAuthMiddleware ইউজারের টোকেন ভ্যালিড কিনা তা পরীক্ষা করে। টোকেন ভুল বা অনুপস্থিত থাকলে ক্লায়েন্ট সরাসরি 401 Unauthorized এরর পায় (ফ্লো হ্যান্ডলার পর্যন্ত যায় না)। টোকেন সঠিক হলে রিকোয়েস্ট হ্যান্ডলারে পাঠানো হয়।

ধাপ ৪: হ্যান্ডলার লেয়ার (rest/handlers/...)
- হ্যান্ডলার হলো HTTP জগতের প্রতিনিধি। এর কাজ শুধু HTTP ইনপুট নেওয়া এবং HTTP আউটপুট দেওয়া।
- এটি r.Body থেকে JSON ডেটা রিড করে ReqCreateProduct স্ট্রাক্টে ডিকোড করে।
- ডেটা সঠিক হলে সেটিকে কোর domain.Product-এ রূপান্তর করে এবং h.svc.Create(...) মেথড কল করে সার্ভিস লেয়ারে পাঠিয়ে দেয়।

ধাপ ৫: সার্ভিস লেয়ার (services/...)
- এটি অ্যাপ্লিকেশনের মূল বিজনেস লজিক লেয়ার। এখানে সব নিয়মকানুন (যেমন: প্রাইস সঠিক কিনা, স্টক আছে কিনা ইত্যাদি) যাচাই করা হয়।
- সার্ভিস লেয়ার ডাটাবেজ চেনে না এবং সরাসরি কোনো SQL লেখে না। সে ডেটা সেভ করার জন্য ডিপেন্ডেন্সি হিসেবে পাওয়া svc.productRepo.Create(...) কল করে দেয়।

ধাপ ৬ ও ৭: রিপোজিটরি ও ডাটাবেজ লেয়ার (repo/ & PostgreSQL)
- রিপোজিটরি লেয়ারে আসল SQL কুয়েরি লেখা থাকে (যেমন: INSERT INTO products ...)।
- এটি dbCon.DB (sqlx ডাটাবেজ হ্যান্ডেল) ব্যবহার করে PostgreSQL সার্ভারে কুয়েরি এক্সিকিউট করে।
- ডাটাবেজে রেকর্ড ইনসার্ট হয়ে নতুন একটি ইউনিক ID তৈরি হয়।

ধাপ ৮, ৯ ও ১০: ফিরতি পথ ও রেসপন্স (Return Flow to Client)
- Repo ডাটাবেজ থেকে পাওয়া নতুন ID-টি মডেলে স্ক্যান করে Service-এ পাঠায়।
- Service কোনো পরিবর্তন বা ভ্যালিডেশন শেষে অবজেক্টটি Handler-এ পাঠায়।
- Handler তখন utils.SendData() ব্যবহার করে Content-Type: application/json এবং HTTP 201 StatusCreated কোড সহ ক্লায়েন্টকে JSON রেসপন্স পাঠিয়ে রিকোয়েস্ট সম্পন্ন করে।

============================================================================================
                               Quick Summary of Core Layers
============================================================================================

Note: /rest, /services, /repo, and /domain are the core architectural layers. 
Other directories (cmd, config, infra, utils) are for bootstrapping, configuration, and utilities.

Forward Flow (Request):
1. /rest      -> Receives HTTP requests, runs middlewares (CORS, JWT auth, logger), matches routes, 
                 extracts & decodes r.Body, maps DTO to the core domain model, and calls /services.
2. /services  -> Core business logic layer. Validates business rules, processes data, 
                 and calls the /repo interface dependency.
3. /repo      -> Data access layer. Executes raw SQL queries against the live database (*sqlx.DB) 
                 and persists/fetches data.

Backward Flow (Response):
4. /repo      -> Scans database records/IDs into the domain model and returns the output (or error) back to /services.
5. /services  -> Receives the output from /repo, performs any post-processing or business validation, 
                 and returns the final result back to /rest.
6. /rest      -> Receives the result from /services, sets the appropriate HTTP status code (201, 400, 500), 
                 sets response headers (application/json), and serializes the JSON response to the client.
*/
