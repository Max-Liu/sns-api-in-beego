package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Top10photo_20141211_153910 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Top10photo_20141211_153910{}
	m.Created = "20141211_153910"
	migration.Register("Top10photo_20141211_153910", m)
}

// Run the migrations
func (m *Top10photo_20141211_153910) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE top10photo(`id` int(11) NOT NULL AUTO_INCREMENT,`photo_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Top10photo_20141211_153910) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	m.Sql("DROP TABLE `top10photo`")
}
