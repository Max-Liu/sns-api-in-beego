package controllers

import "pet/models"

//var fields []string
var sortby []string = []string{"id"}
var order []string = []string{"desc"}

//var query map[string]string = make(map[string]string)
var limit int64 = 10
var offset int64 = 0

func hasMore(query map[string]string, fields []string, offset int64, kind string) int {
	hasMoreId := offset + limit
	switch kind {
	case "likes":
		{
			l, _ := models.GetAllLikes(query, fields, sortby, order, hasMoreId, 1)
			if len(l) == 0 {
				return 0
			} else {
				return 1
			}
		}
	}
	return 0
}
