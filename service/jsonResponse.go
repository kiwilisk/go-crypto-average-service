package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJsonResponse(writer http.ResponseWriter, value interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(value)
	if err != nil {
		err = fmt.Errorf("Failed to encode %v to JSON, %v", value, err)
	}
	return err
}
