package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	fmt.Println("Jai Shree Ram")

	mux := &http.ServeMux{}

	mux.HandleFunc("/greetings/{greeting}", handler)

	urls := []string{
		"/greetings/hello-world",
		"/greetings/good-morning",
		"/greetings/hello-world/extra",
		"/greetings/",
		"/greetings",
		"/messages/hello-world",
	}

	for _, u := range urls {
		req := httptest.NewRequest(http.MethodGet, u, nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		resp := rr.Result()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Request failed: %d %v\n", resp.StatusCode, u)
		}

	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	g := r.PathValue("greeting")
	fmt.Println("Greeting received: %v\n", g)
}
