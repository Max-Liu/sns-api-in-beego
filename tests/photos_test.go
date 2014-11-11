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
	request = helper.MakeUploadRequest(query, testUrl+"/v1/photos/", "POST", "photo", "dog.jpg")
	baseTest(t)
	photoData := jsonData.Data.(map[string]interface{})
	photoId = photoData["Id"].(float64)
}

func TestGetPhotoId(t *testing.T) {
	photoIdStr := strconv.Itoa(int(photoId))
	request, _ = http.NewRequest("GET", testUrl+"/v1/photos/"+photoIdStr, nil)
	baseTest(t)
}
func TestGetMyPhotos(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/photos/?myphoto=1", nil)
	baseTest(t)
}
func TestGetPhotos(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/photos/", nil)
	baseTest(t)
}
func TestGetTimelinePhotos(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/photos/timeline/following", nil)
	baseTest(t)
}

func TestGetPhotoComments(t *testing.T) {
	photoIdStr := strconv.Itoa(int(photoId))
	request, _ = http.NewRequest("GET", testUrl+"/v1/photos?photo_id/"+photoIdStr, nil)
	baseTest(t)
}

func TestGetLikes(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/likes/", nil)
	baseTest(t)
}
func TestCommentPhoto(t *testing.T) {
	query := make(map[string]string)
	photoIdStr := strconv.Itoa(int(photoId))
	query["photo_id"] = photoIdStr
	query["content"] = "testContent afasfaiofaosfoaf"
	request = helper.MakeRequest(query, testUrl+"/v1/comments/", "POST")
	baseTest(t)
}
func TestLikePhoto(t *testing.T) {
	query := make(map[string]string)
	photoIdStr := strconv.Itoa(int(photoId))
	query["photo_id"] = photoIdStr
	request = helper.MakeRequest(query, testUrl+"/v1/likes/", "POST")
	baseTest(t)
}
