package test

import (
	"net/http"
	"pet/utils"
	"testing"
)

func TestGetFollower(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/ul/follower", nil)
	baseTest(t)
}
func TestGetFollowing(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/ul/following", nil)
	baseTest(t)
}
func TestFollowing(t *testing.T) {
	query := make(map[string]string)
	query["following"] = "41"
	request = helper.MakeRequest(query, testUrl+"/v1/ul", "POST")
	baseTest(t)
}

func TestUnFollowing(t *testing.T) {
	query := make(map[string]string)
	query["following"] = "41"
	request = helper.MakeRequest(query, testUrl+"/v1/ul/41", "DELETE")
	baseTest(t)
}
