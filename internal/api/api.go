package api

import (
	"go-file-soap/internal/soap"
	"go-file-soap/internal/storage"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

const maxFileSize = 3 * 1024 * 1024 // 3 Мб

func UploadMTOMHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что заголовок Content-Type = multipart/related
	contentType := r.Header.Get("Content-Type")
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil || !strings.HasPrefix(mediaType, "multipart/related") {
		soap.WriteSOAPResponse(w, "error", "Content-Type must be multipart/related (MTOM)")
		return
	}

	mr := multipart.NewReader(r.Body, params["boundary"])
	var soapEnvelope *soap.SOAPEnvelope
	var filename, fileCID string
	var fileData []byte

	// Парсим файл
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			soap.WriteSOAPResponse(w, "error", "Failed to parse multipart: "+err.Error())
			return
		}
		contentID := part.Header.Get("Content-ID")
		contentType := part.Header.Get("Content-Type")

		if strings.Contains(contentType, "application/xop+xml") {
			envBytes, _ := io.ReadAll(part)
			var env soap.SOAPEnvelope
			if err := soap.UnmarshalEnvelope(envBytes, &env); err != nil {
				soap.WriteSOAPResponse(w, "error", "Invalid SOAP XML: "+err.Error())
				return
			}
			soapEnvelope = &env
			filename = env.Body.UploadFileRequest.Filename
			fileCID = env.Body.UploadFileRequest.File.Href
			fileCID = strings.TrimPrefix(fileCID, "cid:")
		} else if contentID != "" {
			if fileCID != "" && strings.Contains(contentID, fileCID) {
				limitedReader := io.LimitReader(part, maxFileSize+1)
				fileData, err = io.ReadAll(limitedReader)
				if err != nil {
					soap.WriteSOAPResponse(w, "error", "Failed to read file: "+err.Error())
					return
				}
			}
		}
	}

	if soapEnvelope == nil || len(fileData) == 0 {
		soap.WriteSOAPResponse(w, "error", "SOAP envelope or file part missing")
		return
	}
	if int64(len(fileData)) > maxFileSize {
		soap.WriteSOAPResponse(w, "error", "File exceeds 3MB limit")
		return
	}

	if filename == "" {
		soap.WriteSOAPResponse(w, "error", "Filename is required")
		return
	}

	if err := storage.SaveFile(filename, fileData); err != nil {
		soap.WriteSOAPResponse(w, "error", "Failed to save file: "+err.Error())
		return
	}
	soap.WriteSOAPResponse(w, "success", "File saved as "+filename)
}
