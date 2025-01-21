package main

import (
	"fmt"
	"log"
	"net/http"

	"image_uploader/pkg/handler"
)

func main() {
	http.HandleFunc("/upload", handler.UploadHandler)

	fmt.Println("Servidor rodando na porta 8080. Acesse /upload via POST para enviar arquivos.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
