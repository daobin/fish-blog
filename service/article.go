package service

import (
	"github.com/daobin/fish-blog/dao"
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/util"
	"gopkg.in/mgo.v2/bson"
	"math"
	"strings"
	"time"
)

type article struct{}

func (art *article) Add(params model.SaveArticleReq) model.Result[any] {
	// 校验分类是否存在
	cateInfo, err := dao.Category.GetByCateId(params.CateId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if Category.IsExist(cateInfo) == false {
		return model.Error[any](util.CommonError.ToInt(), "分类不存在")
	}

	// 校验标题是否存在
	titleInfo, err := dao.Article.GetByArticleTitle(params.Title, "")
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if titleInfo != nil {
		return model.Error[any](util.CommonError.ToInt(), "标题已存在")
	}

	entity := &model.ArticleEntity{}
	util.Struct2Struct(params, entity)

	entity.Id = bson.NewObjectId()
	entity.ArticleId = util.GenRandomString(8)
	entity.CreatedAt = time.Now().Unix()
	entity.UpdatedAt = entity.CreatedAt

	err = dao.Article.Save(entity)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	return model.Success[any](entity)
}

func (art *article) Update(params model.SaveArticleReq) model.Result[any] {
	// 校验文章是否存在
	info, err := dao.Article.GetByArticleId(params.ArticleId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if info == nil {
		return model.Error[any](util.CommonError.ToInt(), "文章不存在")
	}

	// 校验分类是否存在
	if params.CateId != "" && params.CateId != info.CateId {
		cateInfo, err := dao.Category.GetByCateId(params.CateId)
		if err != nil {
			return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
		}

		if Category.IsExist(cateInfo) == false {
			return model.Error[any](util.CommonError.ToInt(), "分类不存在")
		}

		info.CateId = params.CateId
	}

	// 校验标题是否存在
	if params.Title != "" && params.Title != info.Title {
		titleInfo, err := dao.Article.GetByArticleTitle(params.Title, info.ArticleId)
		if err != nil {
			return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
		}

		if titleInfo != nil {
			return model.Error[any](util.CommonError.ToInt(), "标题已存在")
		}

		info.Title = params.Title
	}

	if params.Description != "" {
		info.Description = params.Description
	}

	if params.State != 0 {
		info.State = params.State
	}
	if params.Sort > 0 {
		info.Sort = params.Sort
	}

	info.UpdatedAt = time.Now().Unix()

	err = dao.Article.Save(info)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	return model.Success[any](info)
}

func (art *article) GetByArticleId(articleId string) model.Result[any] {
	info, err := dao.Article.GetByArticleId(articleId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if info == nil {
		return model.Error[any](util.CommonError.ToInt(), "文章不存在")
	}

	return model.Success[any](info)
}

func (art *article) GetList(params model.GetArticleListReq) model.Result[any] {
	cateIds := strings.Trim(params.CateIds, ", ")
	if cateIds != "" {
		// 如果只存在一个分类，则查询所有的子分类
		if strings.Count(cateIds, ",") == 0 {
			cateIdSlice := Category.getChildrenIds(cateIds)
			params.CateIds = strings.Join(cateIdSlice, ",")
		}
	}

	list, rdCount, err := dao.Article.GetList(&params)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if params.IsReturnPage == false {
		return model.Success[any](list)
	}

	pageCount := int(math.Ceil(float64(rdCount) / float64(params.PageSize)))
	return model.SuccessWithPage[any](list, params.PageIndex, params.PageSize, pageCount, rdCount)
}
