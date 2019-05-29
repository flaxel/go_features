package main

import (
	"encoding/json"
	"fmt"
)

type IOSList struct {
	Message Message `json:"aps"`
}

// omitempty: use it sparingly, it will only be used for Marshalling
type Message struct {
	Alert string `json:"alert,omitempty"`
	Sound string `json:"sound"`
}

func main() {
	// OBJECT
	// m := Message{"test", "default"}
	m := Message{}
	p := &IOSList{m}

	content, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Cannot create json string: ", err)
	}

	fmt.Println(string(content))

	// MAP
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	text, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(text))

	// UNMARSHAL
	b := []byte(`{"aps":{"sound":"sound", "alert":"alert"}}`)
	var p2 IOSList
	err = json.Unmarshal(b, &p2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p2)
}
