package api

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

// GetLocalAppdata
//
//	@Description: 读取localappdata(Windows)
//	@return string
func GetLocalAppdata() string {
	localappdata, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Get LocalAppdata Error: %v", err)
	}
	return localappdata
}

// GetLndConf
//
//	@Description: 读取lnd.conf(Windows)
//	@return string
func GetLndConf() string {
	//localappdata,err:= os.UserCacheDir();
	byteSlice, err := os.ReadFile(GetLocalAppdata() + "\\Lnd\\lnd.conf")
	if err != nil {
		panic(err.Error())
	}
	return string(byteSlice)
}

func DeleteFile(filePath string) bool {
	err := os.RemoveAll(filePath)
	if err != nil {
		log.Printf("File %v delete error: %v", filePath, err)
		return false
	}
	log.Printf("File %v deleted successfully", filePath)
	return true
}

func deleteFilesInLnd() {
	DeleteFile(GetLocalAppdata() + "\\Lnd\\data")
	DeleteFile(GetLocalAppdata() + "\\Lnd\\letsencrypt")
	DeleteFile(GetLocalAppdata() + "\\Lnd\\logs")
	DeleteFile(GetLocalAppdata() + "\\Lnd\\tls.cert")
	DeleteFile(GetLocalAppdata() + "\\Lnd\\tls.key")
}

func StructToJsonString(value any) string {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var str bytes.Buffer
	err = json.Indent(&str, jsonBytes, "", "\t")
	if err != nil {
		log.Fatalf("%v", err)
	}
	result := str.String()
	return result
}
