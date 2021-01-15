package GRequests

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func Unmarshal(data io.Reader,body interface{}){
	bytes, _ := ioutil.ReadAll(data)
	json.Unmarshal(bytes, &body)
}