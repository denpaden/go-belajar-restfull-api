package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetResponseStatus(responseBody map[string]interface{}) string {
	return string(responseBody["status"].(string))
}

func GetResponseCode(responseBody map[string]interface{}) int {
	return int(responseBody["code"].(float64))
}

func ReadResponseBody(response *http.Response) map[string]interface{} {
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	return responseBody
}
