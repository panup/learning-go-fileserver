package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	resp := make(map[string]string)

	resp["fileKey"] = generateKey()
	resp["encryptionKey"] = generateKey()

	if err != nil {
		fmt.Println(err)
		return
	}

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	encryptedBytes, err := encryptBytes(fileBytes, []byte(resp["encryptionKey"]))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("storage/"+resp["fileKey"], encryptedBytes, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonresp, _ := json.Marshal(resp)

	w.Write(jsonresp)
}

func fileDownloadHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Query().Get("filekey")
	encryptionKey := r.URL.Query().Get("encryptionkey")

	if filename == "" || encryptionKey == "" {
		fmt.Println("Filename or key empty from query.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileBytes, err := os.ReadFile("storage/" + filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/octet-stream")

	decryptedBytes, err := decryptBytes(fileBytes, []byte(encryptionKey))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(decryptedBytes)

	err = os.Remove("storage/" + filename)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
