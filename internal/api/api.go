package api

import (
	"encoding/base64"
	"go-file-soap/internal/soap"
	"go-file-soap/internal/storage"
	"io"
	"net/http"
)

func UploadJsonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		soap.WriteSOAPResponse(w, "error", "Failed to read body: "+err.Error())
		return
	}

	var envelope soap.SOAPEnvelope
	if err := soap.UnmarshalEnvelope(bodyBytes, &envelope); err != nil {
		soap.WriteSOAPResponse(w, "error", "Invalid SOAP XML: "+err.Error())
		return
	}
	req := envelope.Body.UploadJSONRequest

	decoded, err := base64.StdEncoding.DecodeString(req.Content)
	if err != nil {
		soap.WriteSOAPResponse(w, "error", "Failed to decode base64 content: "+err.Error())
		return
	}
	if err := storage.SaveFile(req.Filename, decoded); err != nil {
		soap.WriteSOAPResponse(w, "error", "Failed to save file: "+err.Error())
		return
	}
	soap.WriteSOAPResponse(w, "success", "File saved as "+req.Filename)
}
