package GRequests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string, header map[string]string, params interface{}, res interface{}) (*http.Response, error) {
	//创建client
	client := http.Client{}
	//处理data
	bytesData, _ := json.Marshal(params)
	request, _ := http.NewRequest("GET", url, bytes.NewReader(bytesData))
	for key, value := range header {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	if res != nil {
		err = json.NewDecoder(response.Body).Decode(res)
	}
	return response, err
}

func Post(url string, header map[string]string, data interface{}, res interface{}) (*http.Response, error) {
	//创建client
	client := http.Client{}
	request, _ := http.NewRequest("POST", url, nil)
	if data != nil {
		//处理data
		bytesData, _ := json.Marshal(data)
		request, _ = http.NewRequest("POST", url, bytes.NewReader(bytesData))
	}
	for key, value := range header {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	if res != nil {
		// copy body
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		response.Body.Close()
		response.Body = ioutil.NopCloser(buf)
		err = json.NewDecoder(buf).Decode(res)
	}
	return response, err
}

func Put(url string, header map[string]string, data interface{}, res interface{}) (*http.Response, error) {
	//创建client
	client := http.Client{}
	request, _ := http.NewRequest("PUT", url, nil)
	if data != nil {
		//处理data
		bytesData, _ := json.Marshal(data)
		request, _ = http.NewRequest("PUT", url, bytes.NewReader(bytesData))
	}
	for key, value := range header {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	if res != nil {
		// copy body
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		response.Body.Close()
		response.Body = ioutil.NopCloser(buf)
		err = json.NewDecoder(buf).Decode(res)
	}
	return response, err
}
