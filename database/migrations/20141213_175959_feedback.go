package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Feedback_20141213_175959 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Feedback_20141213_175959{}
	m.Created = "20141213_175959"
	migration.Register("Feedback_20141213_175959", m)
}

// Run the migrations
func (m *Feedback_20141213_175959) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE feedback(`id` int(11) NOT NULL AUTO_INCREMENT,`content` varchar(128) NOT NULL,`user_id` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Feedback_20141213_175959) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	m.Sql("DROP TABLE `feedback`")
}
