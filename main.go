package main

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"net/http"
)

var Ctx context.Context = context.TODO()
var DatastoreClient *datastore.Client

func main() {
	dc, err := datastore.NewClient(Ctx, "btc-crawler")

	if err != nil {
		panic(err)
	}

	DatastoreClient = dc

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/signup", SignUpHandler)
	r.HandleFunc("/signin", SignInHandler)
	r.HandleFunc("/stores", Auth(StoresHandler))
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)
}
