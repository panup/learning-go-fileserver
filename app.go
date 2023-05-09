package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		err := os.Mkdir("storage", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/upload", fileUploadHandler)
	mux.HandleFunc("/download", fileDownloadHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("Served at http://localhost%s\n", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
