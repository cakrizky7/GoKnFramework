package models

import (
	"github.com/eaciit/orm"
	"gopkg.in/mgo.v2/bson"
)

func NewUsers() *Users {
	m := new(Users)
	m.Id = bson.NewObjectId()
	return m
}

func (e *Users) RecordID() interface{} {
	return e.Id
}

func (m *Users) TableName() string {
	return "Users"
}

type Users struct {
	orm.ModelBase `bson:"-",json:"-"`
	Id            bson.ObjectId `bson:"_id" , json:"_id" `
	Username      string        `bson:"Username" , json:"Username" `
	Password      string        `bson:"Password" , json:"Password" `
}
