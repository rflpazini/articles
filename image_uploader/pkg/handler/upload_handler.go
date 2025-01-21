package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"image_uploader/internal/service"
)

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	FileName string `json:"file_name,omitempty"`
}

func writeJSONError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{Error: msg})
}

func writeJSONSuccess(w http.ResponseWriter, statusCode int, fileName string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(successResponse{
		FileName: fileName,
	})
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Erro ao ler arquivo do formulário")
		return
	}
	defer file.Close()

	err = service.SaveUploadedFile(file, header.Filename)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Erro ao gravar arquivo: %v", err))
		return
	}

	writeJSONSuccess(w, http.StatusOK, header.Filename)
}
