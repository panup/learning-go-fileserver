# self-learning-go-fileserver
 
 
 Small project as first introduction to Golang.
 Temporal fileserver where uploaded files are encrypted and retrived with filename key and encryption key.
 Once file is retrived it will be deleted from server.
 
 Run with
 
 ```bash
 go run .
 ```
 
 Upload with
 
 ```bash
 curl --location 'http://localhost:8080/upload' \
 --form 'file=@"/location/to/img.jpg"'
 ```
 
 returns
 
 ```json
 {
    "encryptionkey": "a61d3e138ced4c05a07f8c756960b48e",
    "filekey": "5ebc4ea526ca4fecb6558df37f743a6b"
 }
 ```
 
 Download with
 
 ```bash
 curl --location 'http://localhost:8080/download?encryptionkey=a61d3e138ced4c05a07f8c756960b48e&filekey=5ebc4ea526ca4fecb6558df37f743a6b'
 ```
 
 
