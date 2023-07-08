package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ServeData(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	con, err := http.Get("https://www.themealdb.com/api/json/v1/1/random.php")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found")
		return
	}

	defer con.Body.Close()

	data, err := ioutil.ReadAll(con.Body)

	if err != nil {
		w.WriteHeader(http.StatusInsufficientStorage)
		fmt.Fprintf(w, "couldn't decode")
		return
	}
	var dataMap = make(map[string]interface{})

	err = json.Unmarshal(data, &dataMap)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataMap)
}
