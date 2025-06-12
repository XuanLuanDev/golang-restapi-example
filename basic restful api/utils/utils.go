package utils

import(
	"encoding/json"
    "net/http"
)

const ConnectString = "<user>:<password>@tcp(127.0.0.1:3306)/<database-name>"

func ResponseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
    result, _ := json.Marshal(data)
    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(statusCode)
    response.Write(result)
}

func ResponseWithError(response http.ResponseWriter, statusCode int, msg string) {
    ResponseWithJSON(response, statusCode, map[string]string{
        "error": msg,
    })
}