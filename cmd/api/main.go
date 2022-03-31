package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luebken/api-logistics/pkg/model"
)

func main() {
	http.HandleFunc("/vessels", vessels)
	http.HandleFunc("/containers", containers)
	addr := ":8090"
	fmt.Printf("Listen and serve at %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func vessels(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request vessels")
	json.NewEncoder(w).Encode(model.GetVessels())
}
func containers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request containers")
	json.NewEncoder(w).Encode(model.GetContainers())
}
