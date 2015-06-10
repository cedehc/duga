package main

import (
	"fmt"
	"net/http"

	"github.com/coocood/jas"

	"duca/api"
)

func main() {
	db := api.NewFakeDb()
	db.FakeData()

	comics := new(api.ComicId)
	comics.ComicKeeper = db

	router := jas.NewRouter(comics)
	router.BasePath = "/"

	// show paths
	fmt.Println(router.HandledPaths(true))

	http.Handle(router.BasePath, router)
	http.ListenAndServe(":8080", nil)
}
