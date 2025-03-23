package Utility

import (
	MetaData "DistributedFileDBChunkServer/Helper/MetaData"
	"os"
	"path/filepath"
)

func InitiateChunkNameList() {
	err := filepath.Walk(MetaData.FileStoreLocation, func(path string, info os.FileInfo, err error) error {
		if info != nil && info.IsDir() != true {
			MetaData.ChunkNameList = append(MetaData.ChunkNameList, info.Name())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func UpdateChunkNameList(ChunkID string) {
	MetaData.ChunkNameList = append(MetaData.ChunkNameList, ChunkID)
}
