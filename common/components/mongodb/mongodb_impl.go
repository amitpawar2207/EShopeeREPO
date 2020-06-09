package mongodb

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoDriver common
type MongoDriver struct {
	Conn    *mgo.Database
	Session *mgo.Session
	Confg   *MDBConfig
}

//Init Mongo session
func (obj *MongoDriver) Init(conf *MDBConfig) error {
	mongoSession, err := mgo.Dial(conf.URL)
	if err != nil {
		return err
	}
	obj.Session = mongoSession
	obj.Conn = mongoSession.DB(conf.DBName)

	return nil
}

//Insert documents
func (obj *MongoDriver) Insert(collection string, value interface{}) error {
	if err := obj.Conn.C(collection).Insert(value); err != nil {
		return err
	}
	return nil
}

//Close mongoSession
func (obj *MongoDriver) Close(session *MSession) {
	if session == nil {
		return
	}
	session.mgoSession.Close()
}

//Remove documents
func (obj *MongoDriver) Remove(collection string, query map[string]interface{}) error {
	if err := obj.Conn.C(collection).Remove(bson.M(query)); err != nil {
		return err
	}
	return nil
}

//Update documents
func (obj *MongoDriver) Update(collection string, query map[string]interface{}, value interface{}) error {
	if err := obj.Conn.C(collection).Update(bson.M(query), value); err != nil {
		return err
	}
	return nil
}

//FindOne Document
func (obj *MongoDriver) FindOne(collection string, query map[string]interface{}) (ret interface{}, err error) {
	if err := obj.Conn.C(collection).Find(query).One(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//FindAll Document
func (obj *MongoDriver) FindAll(collection string, query map[string]interface{}) (ret []interface{}, err error) {
	if err := obj.Conn.C(collection).Find(query).All(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Count Documents
func (obj *MongoDriver) Count(collection string) (int, error) {
	count, err := obj.Conn.C(collection).Count()
	if err != nil {
		return 0, nil
	}
	return count, nil
}

//Find and return result
func (obj *MongoDriver) Find(collection string, query map[string]interface{}, selectField map[string]interface{}, skip, limit int) (ret []interface{}, err error) {
	err = obj.Conn.C(collection).Find(query).Select(bson.M(selectField)).Skip(skip).Limit(limit).All(&ret)
	if err != nil {
		return nil, fmt.Errorf("Error while finding data")
	}
	return ret, nil
}
