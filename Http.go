package GRequests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Get(url string,header map[string]string,params map[string]interface{})*http.Response{
	//创建client
	client := http.Client{}
	//处理data
	bytesData, _ := json.Marshal(params)
	request, _ := http.NewRequest("GET", url,bytes.NewReader(bytesData))
	for key,value :=range header{
		request.Header.Set(key,value)
	}
	response, _ := client.Do(request)
	return response
}

func Post(url string,header map[string]string,data map[string]interface{})*http.Response{
	//创建client
	client := http.Client{}
	//处理data
	bytesData, _ := json.Marshal(data)
	request, _ := http.NewRequest("GET", url,bytes.NewReader(bytesData))
	for key,value :=range header{
		request.Header.Set(key,value)
	}
	response, _ := client.Do(request)
	return response
}
