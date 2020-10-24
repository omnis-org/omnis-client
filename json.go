package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Output structure to gather machine infos
type Output struct {
	OS       string `json:"os"`
	HostName string `json:"hostname"`
	Platform string `json:"platform"`
}

func main() {
	out := Output{
		OS:       "pulsz",
		HostName: "yolol",
		Platform: "salutcava",
	}
	jsonstr, err := json.Marshal(out)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	f, err := os.OpenFile("cc.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Errorf("%v", err)
	}

	_, err = f.Write(jsonstr)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
