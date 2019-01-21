package dao

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"one/go-misc/modules/map/model"
	"one/go-misc/utils"
)
var dbName =  utils.MongoSetting.DbName
const collectionName = "testgo"

func Insert() {
	var (
		err             error
		collection      *mongo.Collection
		insertOneRes    *mongo.InsertOneResult
	)
	//链接mongo服务
	collection = getMongoCollection(dbName,collectionName)

	test := model.Test{
		Name:       "test",
		Pwd:        "123",
		Age:        10,
		CreateTime: 1,
	}
	//插入一条数据
	if insertOneRes, err = collection.InsertOne(getContext(), test);
	err != nil {
		checkErr(err)
	}
	fmt.Printf("InsertOne插入的消息ID:%v\n", insertOneRes.InsertedID)
}
