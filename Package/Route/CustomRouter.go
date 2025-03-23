package Route

import (
	"DistributedFileDBChunkServer/Helper/RouterURL"
	"DistributedFileDBChunkServer/Package/Controller"

	"github.com/gorilla/mux"
)

func CustomRouter(Mux *mux.Router) {

	Mux.HandleFunc(RouterURL.Return, Controller.ReturnData).Methods("GET")
	Mux.HandleFunc(RouterURL.Update, Controller.UpdateData).Methods("PUT")
	Mux.HandleFunc(RouterURL.Add, Controller.AddData).Methods("POST")

}
