package main

import (
	"DistributedFileDBChunkServer/Package/Route"
	"DistributedFileDBChunkServer/Package/Utility"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Utility.InitiateChunkNameList()

	//Utility.InitiateSocketConnection()

	go Utility.InitiateGRPCConnection()

	//go Utility.Listener()

	muxRouter := mux.NewRouter()

	Route.CustomRouter(muxRouter)

	http.Handle("/", muxRouter)

	http.ListenAndServe("localhost:9030", muxRouter)

	defer Utility.TerminateSocketConnection()

}
