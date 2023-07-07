package main

import (
	"encoding/json"
	"fmt"
	"http-request-send/utils"
	"io"
	"net/http"
)

type champion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var champions []*champion

func main() {
	resp, err := http.Get("http://localhost:4000")
	utils.HandleErr(err)

	jsonBody, err := io.ReadAll(resp.Body)
	utils.HandleErr(err)
	utils.HandleErr(json.Unmarshal(jsonBody, &champions))

	for _, champion := range champions {
		fmt.Printf("ID %d Name %s \n", champion.ID, champion.Name)
	}
}
