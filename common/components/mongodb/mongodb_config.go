package mongodb

import "gopkg.in/mgo.v2"

//MDBConfig is the configueartion for mongoDB
type MDBConfig struct {
	URL    string
	DBName string
}

//MSession is the mongo Session
type MSession struct {
	mgoSession *mgo.Session
}
