package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	sendData()
}

func handleCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func sendData() {
	http.HandleFunc("/", serveData)
	http.ListenAndServe(":8080", nil)
}

func serveData(w http.ResponseWriter, r *http.Request) {
	handleCors(&w)
	var urlFetch, fetchErr = http.Get("https://www.themealdb.com/api/json/v1/1/random.php")

	if fetchErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"err": "cannot Fetch data",
		})
	}
	defer urlFetch.Body.Close()

	var databyte, decodeErr = io.ReadAll(urlFetch.Body)
	if decodeErr != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"err": "couldnt decode bytes",
		})
	}

	var data map[string]interface{} = map[string]interface{}{}
	json.Unmarshal(databyte, &data)
	json.NewEncoder(w).Encode(data)
}
