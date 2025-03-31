package main

import (
	"DistributedFileDBChunkServer/Package/Route"
	"DistributedFileDBChunkServer/Package/Utility"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Utility.InitiateChunkNameList()

	Utility.InitiateSocketConnection()

	muxRouter := mux.NewRouter()

	Route.CustomRouter(muxRouter)

	http.Handle("/", muxRouter)

	http.ListenAndServe("localhost:9030", muxRouter)
}
