package tool

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out.String())
}
