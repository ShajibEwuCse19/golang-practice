package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Logger is an HTTP middleware that measures and logs the HTTP method, path, and execution duration of each request.
// It wraps the next http.Handler in the chain and invokes next.ServeHTTP(w, r).
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			end := time.Now()
			timeDiff := end.Sub(start)
			fmt.Printf("Logger: %s %s %s\n", r.Method, r.URL.Path, timeDiff.String())
		},
	)
}

/*
=> Logger()-এর return type http.Handler, কিন্তু আমরা return করছি http.HandlerFunc। তাহলে type mismatch হচ্ছে না কেন?
Short ans: http.HandlerFunc implement http.Handler interface, ServeHTTP. So, http.HandlerFunc == http.Handler.
Answer: http.HandlerFunc হলো এমন একটি named function type, যা http.Handler interface implement করে।
তাই http.HandlerFunc-এর value-কে http.Handler হিসেবে return করা যায়।

=> http.Handler আসলে কী?
Go-এর net/http package-এ http.Handler হলো একটি interface।

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
এই interface বলছে: যে type-এর ServeHTTP(http.ResponseWriter, *http.Request) method আছে [same signature],
সেই type http.Handler interface implement করতে পারে।

Example:
type MyHandler struct{}

func (h MyHandler) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request,
) {
    fmt.Fprintln(w, "Hello")
}
এখন MyHandler-এর ServeHTTP method আছে। তাই MyHandler http.Handler interface implement করে। So, MyHandler == http.Handler.

=> http.HandlerFunc কী?
net/http package-এ এর definition:
type HandlerFunc func(ResponseWriter, *Request)

এটি একটি named function type। অর্থাৎ, একটি নির্দিষ্ট signature-এর function-কে Go একটি named type হিসেবে প্রকাশ করছে।
কিন্তু শুধু function type হলেই সেটি http.Handler implement করে না। HandlerFunc-এর ওপর একটি method-ও define করা আছে:

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}

খেয়াল করো:
- f হলো একটি HandlerFunc value।
- f(w, r) দিয়ে সেই function-টিকে সরাসরি call করা হচ্ছে।
- ServeHTTP() method-এর কারণে HandlerFunc এখন http.Handler interface-এর contract পূরণ করছে।
অর্থাৎ, একটি function-কে এমনভাবে wrap করা হয়েছে, যাতে সেটিকে http.Handler হিসেবে ব্যবহার করা যায়।

Summary:
http.Handler = একটি interface, যার একটি method ServeHTTP() আছে।
http.HandlerFunc = একটি function type, যা http.Handler interface implement করে।
Logger() → http.Handler return করে (handler → handler → handler → handler)
হিসেবে।
*/
