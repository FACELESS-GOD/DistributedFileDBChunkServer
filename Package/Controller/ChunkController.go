package Controller

import (
	"DistributedFileDBChunkServer/Helper/MetaData"
	"DistributedFileDBChunkServer/Helper/RouterURL"
	"DistributedFileDBChunkServer/Helper/StructStore"
	"DistributedFileDBChunkServer/Package/Utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ReturnData(writer http.ResponseWriter, Req *http.Request) {
	// Response Object
	Response := StructStore.GenericResponseData{}
	//--
	// Getting the ChunkID from the Query Parameters
	params := mux.Vars(Req)
	ChunkID := params[RouterURL.QueryParam]
	//--

	// Return Data
	filepath := MetaData.FileStoreLocation + "/" + ChunkID
	data, err := os.ReadFile(filepath)
	if err != nil {
		InvalidOperationResponse2(writer, Response, "SignUp was Unsuccessfull")
		return
	}

	Response.Data = append(Response.Data, data...)

	ValidOperationResponse2(writer, Response, "SignUp was successfull")
	return
}

func UpdateData(writer http.ResponseWriter, Req *http.Request) {
	// Response Object
	Response := StructStore.GenericResponseMessage{}
	//--

	// Getting the ChunkID from the Query Parameters
	params := mux.Vars(Req)
	ChunkID := params[RouterURL.QueryParam]
	//--

	// Getting File from Request

	Req.ParseMultipartForm(1000 << 20)
	file, handler, err := Req.FormFile("ChunkFile")

	if err != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}
	defer file.Close()
	fmt.Print(handler.Size)
	//--

	// Processing the File

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}
	fileLoc := MetaData.FileStoreLocation + "/" + ChunkID
	WritErr := os.WriteFile(fileLoc, fileBytes, 0666)

	if WritErr != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}

	//--

	ValidOperationResponse(writer, Response, "SignUp was successfull")
	return
}

func AddData(writer http.ResponseWriter, Req *http.Request) {
	// Response Object
	Response := StructStore.GenericResponseMessage{}
	//--

	// Getting the ChunkID from the Query Parameters
	params := mux.Vars(Req)
	ChunkID := params[RouterURL.QueryParam]
	//--

	// Getting File from Request

	Req.ParseMultipartForm(1000 << 20)
	file, handler, err := Req.FormFile("ChunkFile")

	if err != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}
	defer file.Close()
	fmt.Print(handler.Size)
	//--

	// Processing the File
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}

	fileLoc := MetaData.FileStoreLocation + "/" + ChunkID

	//file, er := os.Create(fileLoc)

	// if er != nil {
	//     panic(err)
	// }

	WritErr := os.WriteFile(fileLoc, fileBytes, 0666)

	if WritErr != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}

	//--

	go Utility.UpdateChunkNameList(ChunkID)

	ValidOperationResponse(writer, Response, "SignUp was successfull")
	return

}

func InvalidOperationResponse(writer http.ResponseWriter, Response StructStore.GenericResponseMessage, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(res)
}

func ValidOperationResponse(writer http.ResponseWriter, Response StructStore.GenericResponseMessage, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func InvalidOperationResponse2(writer http.ResponseWriter, Response StructStore.GenericResponseData, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(res)
}

func ValidOperationResponse2(writer http.ResponseWriter, Response StructStore.GenericResponseData, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
