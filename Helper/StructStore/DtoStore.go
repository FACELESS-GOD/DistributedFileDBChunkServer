package StructStore

type GenericResponseMessage struct {
	Message string
}

type GenericResponseData struct {
	Data    []byte
	Message string
}

type ChunkMapping struct {
	ServerID       string
	AvailableSpace int64
	ChunkList      []string
}
