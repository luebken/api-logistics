package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	idParam := req.URL.Query().Get("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	if id != 0 {
		fmt.Printf("request vessel for id %d\n", id)
		v := model.GetVessel(id)
		if v.ID != 0 {
			json.NewEncoder(w).Encode(v)
		} else {
			json.NewEncoder(w).Encode(nil)
		}
	} else {
		fmt.Println("request all vessels")
		json.NewEncoder(w).Encode(model.GetAllVessels())
	}
}
func containers(w http.ResponseWriter, req *http.Request) {
	idParam := req.URL.Query().Get("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	if id != 0 {
		fmt.Printf("request containers for id %d\n", id)
		c := model.GetContainer(id)
		if c.ID != 0 {
			json.NewEncoder(w).Encode(c)
		} else {
			json.NewEncoder(w).Encode(nil)
		}
	} else {
		fmt.Println("request all containers")
		json.NewEncoder(w).Encode(model.GetAllContainers())
	}
}
