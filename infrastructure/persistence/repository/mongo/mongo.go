package mongo

import (
	"ddd-test/infrastructure/config"
	"gopkg.in/mgo.v2"
)

var MgoDB *mgo.Session

func StartMongo(addr string) {
	var err error
	MgoDB, err = mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	MgoDB.SetMode(mgo.Strong, true)
}

func GetCol(db *mgo.Session, col string) *mgo.Collection {
	return db.DB(config.Cfg.DBName).C(col)
}
