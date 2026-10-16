package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"ecom-project-improved/middleware"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body. Give me valid json", 400)
		return
	}

	newProduct.ID = uint8(len(productList) + 1)
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) { // data interface{} - it can take any type of data (int, float, struct, list etc).
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, "Product created successfully")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	const port = ":8080"
	freePort(port) // Automatically check and kill any old or suspended process holding this port

	mux := http.NewServeMux() // router (multiplexer)

	// here, mux.Handle expects a pattern and a http.Handler. So, I have created a middleware for handle CORS error and OPTION requests
	// The, HandleCorsMiddleware checks all CORS related errors and finally return a http.Handler. So I passed it to mux.Handle. This is how middleware works in golang.
	mux.Handle("GET /products", middleware.HandleCorsMiddleware(http.HandlerFunc(getProducts))) // route (endpoint) and request handler with explicit HTTP method (GET). 
	mux.Handle("POST /create-product", middleware.HandleCorsMiddleware(http.HandlerFunc(createProduct))) // route (endpoint) and request handler with explicit HTTP method (POST).

	fmt.Println("Server running on port: 8080")

	err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests. 

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

type Product struct {
	ID          uint8   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgURL      string  `json:"imgURL"`
}

var productList []Product

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Good condition sweet orange",
		Price:       50.60,
		ImgURL:      "https://www.oragneimage.com/imgOrange.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Good condition sweet Apple",
		Price:       150.20,
		ImgURL:      "https://www.appleimage.com/imgApple.png",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Grape",
		Description: "Good condition sweet Grape",
		Price:       90.60,
		ImgURL:      "https://www.grapeimage.com/imgGrape.png",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Good condition sweet Banana",
		Price:       30.60,
		ImgURL:      "https://www.bananaimage.com/imgBanana.png",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Good condition sweet Mango",
		Price:       120.60,
		ImgURL:      "https://www.mangoimage.com/imgMango.png",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}

/* ==================FREE-PORT-LOGIC====================== */

// freePort finds and terminates any existing process running on the specified port
// so the server can always start on the same port without "address already in use" errors.
func freePort(port string) {
	// Remove leading colon if present (e.g. ":8080" becomes "8080")
	portNum := strings.TrimPrefix(port, ":")

	// Prepare shell command "lsof -t -i:8080" to look up process IDs (PIDs) using this port
	// -t: terse output (prints only process IDs)
	// -i: selects network socket matching the port
	cmd := exec.Command("lsof", "-t", "-i:"+portNum)

	// Run command and capture output (contains PID list as text)
	output, err := cmd.Output()
	if err != nil {
		// If command returns error or no process is found, return safely
		return
	}

	// Get current process ID so we don't accidentally terminate our own program
	currentPID := os.Getpid()

	// Split the output string into individual PID strings (handles newlines/spaces)
	pids := strings.Fields(string(output))

	// Loop through each process ID found on this port
	for _, pidStr := range pids {
		// Convert the PID from string to integer
		pid, err := strconv.Atoi(pidStr)

		// Make sure conversion succeeded and that the PID is not our own process
		if err == nil && pid != currentPID {
			// Locate the OS process by its PID
			if proc, err := os.FindProcess(pid); err == nil {
				// Print information about the process being terminated
				fmt.Printf("Freeing port %s (killing old process PID %d)...\n", portNum, pid)

				// Terminate the process (sends SIGKILL)
				_ = proc.Kill()

				// Wait briefly (200ms) to allow the OS network stack to release the socket
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}
