package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id,string"`
}

type helloWroldRequest struct {
	Name string `json:"name"`
}

type validationContextKey string

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWroldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(w, r)
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: fmt.Sprintf("Hellow %s", name)}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func main() {
	port := 7000

	// http.HandleFunc("/helloworld", helloWorldHandler)
	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)
	imageHandler := http.FileServer(http.Dir("./images"))
	http.Handle("/image/", http.StripPrefix("/image/", imageHandler))

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

// func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	var request helloWroldRequest
// 	err = json.Unmarshal(body, &request)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	response := helloWorldResponse{Message: fmt.Sprintf("Hellow %s", request.Name)}
// 	data, err := json.Marshal(response)
// 	if err != nil {
// 		panic("Ooops")
// 	}
// 	fmt.Fprintf(w, string(data))
// }

// func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 	var request helloWroldRequest
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&request)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	response := helloWorldResponse{Message: fmt.Sprintf("Hellow %s", request.Name)}
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(&response)
// }

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
