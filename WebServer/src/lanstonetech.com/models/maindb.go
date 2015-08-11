package models

import (
	// "github.com/astaxie/beego"
	// "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"lanstonetech.com/common"
	"strconv"
	"time"
)

func init() {
	//Register model
	orm.RegisterModel(new(StoneUser), new(StoneCategory), new(StoneTopic), new(StoneTopicreply))

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//set default database
	orm.RegisterDataBase("default", "mysql", "root:lanstonetech@/stone?charset=utf8", 30)

	//orm.Debug = true
	//new tables
	orm.RunSyncdb("default", false, false)
}

func VerifyUser(account, password string) (bool, error) {
	ORM := orm.NewOrm()
	//var users []StoneUser
	users := make([]StoneUser, 0)
	_, err := ORM.Raw("SELECT * from stone_user").QueryRows(&users)
	if err != nil {
		return false, nil
	}

	for _, user := range users {
		if user.Account == account && user.Password == password {
			return true, nil
		}
	}

	return false, nil
}

func AddUser(user StoneUser) error {
	ORM := orm.NewOrm()

	user.Account = common.MakeMD5(user.Telphone)
	user.Password = common.MakeMD5(user.Password)

	_, err := ORM.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserPassword(telphone, email, new_password string) (bool, error) {
	ORM := orm.NewOrm()
	user := new(StoneUser)

	err := ORM.Raw("SELECT * from stone_user where telphone=? and email=?", telphone, email).QueryRow(user)
	if err != nil {
		return false, err
	}

	user.Password = common.MakeMD5(new_password)
	_, err = ORM.Update(user)
	if err != nil {
		return false, err
	}

	return true, err
}

func GetUserName(account string) (string, error) {
	ORM := orm.NewOrm()

	//md5_account := common.MakeMD5(account)
	// u := StoneUser{Account: md5_account}
	// err := ORM.Read(&u)

	var name string
	err := ORM.Raw("SELECT name FROM stone_user WHERE account = ?", account).QueryRow(&name)
	if err != nil {
	}

	return name, err
}

func AddCategory(name string) error {
	ORM := orm.NewOrm()

	cate := &StoneCategory{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := ORM.QueryTable("stone_category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = ORM.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllCategories() ([]*StoneCategory, error) {
	ORM := orm.NewOrm()

	cates := make([]*StoneCategory, 0)

	qs := ORM.QueryTable("stone_category")
	_, err := qs.All(&cates)

	return cates, err
}

func DelCategory(id string) error {
	ORM := orm.NewOrm()

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	cate := &StoneCategory{Id: cid}
	_, err = ORM.Delete(cate)

	return err
}

func AddTopic(account, title, content string) error {
	ORM := orm.NewOrm()

	topic := &StoneTopic{
		Account:   account,
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		Views:     0,
		ReplyTime: time.Now(),
	}

	_, err := ORM.Insert(topic)

	return err
}

func GetAllTopics(isDesc bool, account string) ([]*StoneTopic, error) {
	ORM := orm.NewOrm()
	topics := make([]*StoneTopic, 0)

	if isDesc {
		//_, err = qs.OrderBy("-created").All(&topics)
		_, err := ORM.Raw("SELECT * from stone_topic order by id desc").QueryRows(&topics)
		if err != nil {
			return nil, err
		}
	} else {
		// _, err = qs.All(&topics)
		_, err := ORM.Raw("SELECT * from stone_topic order by id").QueryRows(&topics)
		if err != nil {
			return nil, err
		}
	}

	for i, _ := range topics {
		if topics[i].Account == account {
			topics[i].Ishost = true
		} else {
			topics[i].Ishost = false
		}
	}

	return topics, nil
}

func GetTopicByID(tid string) (*StoneTopic, error) {
	ORM := orm.NewOrm()

	tidnum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	topic := new(StoneTopic)

	qs := ORM.QueryTable("stone_topic")
	err = qs.Filter("id", tidnum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = ORM.Update(topic)

	return topic, err
}

func ModifyTopic(account, tid, title, content string) error {
	tidnum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	ORM := orm.NewOrm()
	topic := &StoneTopic{Id: tidnum, Account: account}
	if ORM.Read(topic) == nil {
		topic.Account = account
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		ORM.Update(topic)

		return nil
	}

	return err
}

func DeleteTopic(tid string) error {
	tidnum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	ORM := orm.NewOrm()
	topic := &StoneTopic{Id: tidnum}
	_, err = ORM.Delete(topic)

	return err
}

func AddTopicReply(tid string, reply_id int64, account, name, content string) error {
	tidnum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	ORM := orm.NewOrm()

	reply := &StoneTopicreply{
		TopicId:   tidnum,
		ReplyId:   reply_id,
		Account:   account,
		Name:      name,
		Content:   content,
		ReplyTime: time.Now(),
	}

	{
		_, err = ORM.Insert(reply)

		topic := new(StoneTopic)
		qs := ORM.QueryTable("stone_topic")
		err = qs.Filter("id", tidnum).One(topic)
		if err != nil {
			return err
		}

		topic.ReplyCount++
		_, err = ORM.Update(topic)
	}

	return err
}

func DelTopicReply(reply_id string) error {
	_reply_id, err := strconv.ParseInt(reply_id, 10, 64)
	if err != nil {
		return err
	}

	ORM := orm.NewOrm()
	reply := &StoneTopicreply{Id: _reply_id}
	_, err = ORM.Delete(reply)

	return err
}

func GetTopicReplys(tid string) ([]*StoneTopicreply, error) {
	ORM := orm.NewOrm()

	replys := make([]*StoneTopicreply, 0)
	_, err := ORM.Raw("SELECT * from stone_topicreply where topic_id=? order by id", tid).QueryRows(&replys)
	if err != nil {
		return nil, err
	}

	return replys, nil
}
