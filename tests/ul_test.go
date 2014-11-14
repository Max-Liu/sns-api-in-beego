package test

import (
	"net/http"
	"pet/utils"
	"testing"
)

func TestGetFollower(t *testing.T) {
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/ul/follower", nil)
	client.BaseTest(t)
}
func TestGetFollowing(t *testing.T) {
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/ul/following", nil)
	client.BaseTest(t)
}
func TestFollowing(t *testing.T) {
	query := make(map[string]string)
	query["following"] = "41"
	client.Request = helper.MakeRequest(query, testUrl+"/v1/ul", "POST")
	client.BaseTest(t)
}

func TestUnFollowing(t *testing.T) {
	query := make(map[string]string)
	query["following"] = "41"
	client.Request = helper.MakeRequest(query, testUrl+"/v1/ul/41", "DELETE")
	client.BaseTest(t)
}
