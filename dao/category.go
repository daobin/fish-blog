package dao

import (
	"errors"
	"github.com/daobin/fish-blog/conf"
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/util"
	"github.com/daobin/gotools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const tableCategory = "fb_category"

type category struct{}

func (cate *category) Save(entity *model.CategoryEntity) error {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	_, err = conn.C(tableCategory).UpsertId(entity.Id, entity)
	if err != nil {
		log.Println("mongo更新（添加）错误：", err.Error())
		return util.DbAbnormal.NewError()
	}

	return nil
}

func (cate *category) GetByCateId(cateId string) (*model.CategoryEntity, error) {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	info := &model.CategoryEntity{}
	err = conn.C(tableCategory).Find(bson.M{"cateId": cateId}).One(info)
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			return nil, nil
		}

		log.Println("mongo查询错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}

	return info, nil
}

func (cate *category) GetByCateName(cateName, filterOutCateId string) (*model.CategoryEntity, error) {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	info := &model.CategoryEntity{}
	err = conn.C(tableCategory).Find(bson.M{"name": cateName, "cateId": bson.M{"$ne": filterOutCateId}}).One(info)
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			return nil, nil
		}

		log.Println("mongo查询错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}

	return info, nil
}

func (cate *category) GetList(params *model.GetCategoryListReq) ([]model.CategoryEntity, int, error) {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return nil, 0, util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	if params.State == util.StateDelete {
		return []model.CategoryEntity{}, 0, nil
	}

	// 查询条件
	find := bson.M{"state": bson.M{"$ne": util.StateDelete}}
	if params.Name != "" {
		find["name"] = bson.M{"$regex": params.Name}
	}
	if params.ParentId != "" {
		find["parentId"] = params.ParentId
	}
	if params.State != 0 {
		find["state"] = params.State
	}

	if params.PageIndex < 1 {
		params.PageIndex = 1
	}
	if params.PageSize < 1 {
		params.PageSize = 10
	}

	// 获取数据总量
	rdCount, err := conn.C(tableCategory).Find(find).Count()
	if err != nil {
		log.Println("mongo查询错误：", err.Error())
		return nil, 0, util.DbAbnormal.NewError()
	}

	list := make([]model.CategoryEntity, 0)
	if params.IsReturnPage == false {
		err = conn.C(tableCategory).Find(find).Sort("state", "sort").All(&list)
		if err != nil {
			log.Println("mongo查询错误：", err.Error())
			return nil, 0, util.DbAbnormal.NewError()
		}

		return list, rdCount, nil
	}

	// 获取分页数据
	skip := (params.PageIndex - 1) * params.PageSize
	err = conn.C(tableCategory).Find(find).Sort("state", "sort").Skip(skip).Limit(params.PageSize).All(&list)
	if err != nil {
		log.Println("mongo查询错误：", err.Error())
		return nil, 0, util.DbAbnormal.NewError()
	}

	return list, rdCount, nil
}
