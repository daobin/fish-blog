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
	"strings"
)

const tableArticle = "fb_article"

type article struct{}

func (art *article) Save(entity *model.ArticleEntity) error {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	_, err = conn.C(tableArticle).UpsertId(entity.Id, entity)
	if err != nil {
		log.Println("mongo更新（添加）错误：", err.Error())
		return util.DbAbnormal.NewError()
	}

	return nil
}

func (art *article) GetByArticleId(articleId string) (*model.ArticleEntity, error) {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	info := &model.ArticleEntity{}
	err = conn.C(tableArticle).Find(bson.M{"articleId": articleId}).One(info)
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			return nil, nil
		}

		log.Println("mongo查询错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}

	return info, nil
}

func (art *article) GetByArticleTitle(title, filterOutArticleId string) (*model.ArticleEntity, error) {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	info := &model.ArticleEntity{}
	err = conn.C(tableArticle).Find(bson.M{"title": title, "articleId": bson.M{"$ne": filterOutArticleId}}).One(info)
	if err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			return nil, nil
		}

		log.Println("mongo查询错误：", err.Error())
		return nil, util.DbAbnormal.NewError()
	}

	return info, nil
}

func (art *article) GetList(params *model.GetArticleListReq) ([]model.ArticleEntity, int, error) {
	conn, err := gotools.DB.Mongo.Get(conf.App.GetString("dbTag.mongo"))
	if err != nil {
		log.Println("mongo连接错误：", err.Error())
		return nil, 0, util.DbAbnormal.NewError()
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	// 查询条件
	find := bson.M{}

	cateIds := strings.Trim(params.CateIds, ", ")
	if cateIds != "" {
		find["cateId"] = bson.M{"$in": strings.Split(cateIds, ",")}
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
	rdCount, err := conn.C(tableArticle).Find(find).Count()
	if err != nil {
		log.Println("mongo查询错误：", err.Error())
		return nil, 0, util.DbAbnormal.NewError()
	}

	list := make([]model.ArticleEntity, 0)
	if params.IsReturnPage == false {
		err = conn.C(tableArticle).Find(find).Sort("state", "sort", "-createdAt").All(&list)
		if err != nil {
			log.Println("mongo查询错误：", err.Error())
			return nil, 0, util.DbAbnormal.NewError()
		}

		return list, rdCount, nil
	}

	// 获取分页数据
	skip := (params.PageIndex - 1) * params.PageSize
	err = conn.C(tableArticle).Find(find).Sort("state", "sort").Skip(skip).Limit(params.PageSize).All(&list)
	if err != nil {
		log.Println("mongo查询错误：", err.Error())
		return nil, 0, util.DbAbnormal.NewError()
	}

	return list, rdCount, nil
}
