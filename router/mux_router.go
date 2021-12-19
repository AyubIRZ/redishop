package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

var (
	r = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) SERVE(port string) error {
	fmt.Printf("Mux HTTP server running on port %v \n", port)
	err := http.ListenAndServe(port, r)

	return err
}

func (*muxRouter) GetVars(r *http.Request) map[string]string {
	params := mux.Vars(r)

	return params
}
