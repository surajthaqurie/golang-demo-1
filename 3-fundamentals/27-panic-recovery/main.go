package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func sillySusan() {
	fmt.Println("Silly susan called")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Hacker have been thwarted")
		}
	}()
	panickingPeter()
	fmt.Println("silly susan finished")
}

func panickingPeter() {
	fmt.Println("Panicking peter called")

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Hacker have been thwarted")
	// 	}
	// }()

	panic("oh no")
	fmt.Println("Panicking peter finished")

}

func main() {
	fmt.Println("Panic! in the GO APP")

	// defer func() {
	// 	fmt.Println("Im a deferred function")
	// }()

	// panic("oh no")
	sillySusan()

	fmt.Println("End of main")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("Unexpected error!")
	})

	http.ListenAndServe(":1123", handlers.RecoveryHandler()(r))
}

/*
type recoveryHandler struct {
	handler http.Handler
}

func (h *recoveryHandler) log(err interface{}) {
	log.Printf("Recovered from panic: %v\n", err)
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.log(err)
		}
	}()

	h.handler.ServeHTTP(w, req)
}
*/
