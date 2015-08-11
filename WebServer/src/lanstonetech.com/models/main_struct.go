package models

import (
	"time"
)

type StoneUser struct {
	Id       int `orm:"auto"`
	Account  string
	Name     string `orm:"size(100)"`
	Email    string
	Telphone string `orm:"size(11)"`
	Password string
}

// 分类
type StoneCategory struct {
	Id              int64 `orm:"auto"`
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserID int64
}

// 文章
type StoneTopic struct {
	Id              int64 `orm:"auto"`
	Uid             int64
	Account         string
	Title           string
	Content         string `orm:"size(10000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserID int64
	Ishost          bool
}

// 文章
type StoneTopicreply struct {
	Id        int64 `orm:"auto"`
	TopicId   int64
	ReplyId   int64
	Account   string
	Name      string
	Content   string `orm:"size(10000)"`
	ReplyTime time.Time
}
