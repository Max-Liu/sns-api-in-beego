package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func baseTest(t *testing.T) {
	resp, err := client.Do(request)
	if err != nil {
		t.Error(err.Error())
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	if resp.StatusCode != 200 {
		//spew.Dump(body.Bytes())
		log.Println("Printing error html file in current Path:output.html")
		ioutil.WriteFile("output.html", bodyByte, 0644)
		t.Error(resp.StatusCode)
		t.FailNow()
	} else {
		t.Log("status 200")
	}
	err = json.Unmarshal(bodyByte, &jsonData)
	if err != nil {
		t.Error(err.Error())
	}
	if jsonData.Err != 0 {
		t.Error(jsonData.Msg)
	}
}
