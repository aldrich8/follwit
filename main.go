package main

import (
	"follwit/routers"
	"net/http"
)

func main() {
	router := routers.SetupRoute()

	err := http.ListenAndServe(":5200", router)

	if err != nil {
		panic(err)
	}
}
