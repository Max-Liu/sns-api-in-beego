package test

import (
	"net/http"
	"pet/utils"
	"strconv"
	"testing"
)

var photoId float64

func TestCreatePhoto(t *testing.T) {
	query := make(map[string]string)
	query["title"] = "mydog"
	client.Request = helper.MakeUploadRequest(query, testUrl+"/v1/photos/", "POST", "photo", "dog.jpg")
	client.BaseTest(t)
	photoData := jsonData.Data.(map[string]interface{})
	photoId = photoData["Id"].(float64)
}

func TestGetPhotoId(t *testing.T) {
	photoIdStr := strconv.Itoa(int(photoId))
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/photos/"+photoIdStr, nil)
	client.BaseTest(t)
}
func TestGetMyPhotos(t *testing.T) {
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/photos/?myphoto=1", nil)
	client.BaseTest(t)
}
func TestGetPhotos(t *testing.T) {
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/photos/", nil)
	client.BaseTest(t)
}
func TestGetTimelinePhotos(t *testing.T) {
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/photos/timeline/following", nil)
	client.BaseTest(t)
}

func TestGetPhotoComments(t *testing.T) {
	photoIdStr := strconv.Itoa(int(photoId))
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/photos?photo_id/"+photoIdStr, nil)
	client.BaseTest(t)
}

func TestGetLikes(t *testing.T) {
	client.Request, _ = http.NewRequest("GET", testUrl+"/v1/likes/", nil)
	client.BaseTest(t)
}
func TestCommentPhoto(t *testing.T) {
	query := make(map[string]string)
	photoIdStr := strconv.Itoa(int(photoId))
	query["photo_id"] = photoIdStr
	query["content"] = "testContent afasfaiofaosfoaf"
	client.Request = helper.MakeRequest(query, testUrl+"/v1/comments/", "POST")
	client.BaseTest(t)
}
func TestLikePhoto(t *testing.T) {
	query := make(map[string]string)
	photoIdStr := strconv.Itoa(int(photoId))
	query["photo_id"] = photoIdStr
	client.Request = helper.MakeRequest(query, testUrl+"/v1/likes/", "POST")
	client.BaseTest(t)
}
