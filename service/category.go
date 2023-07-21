package service

import (
	"github.com/daobin/fish-blog/dao"
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/util"
	"math"
)

type category struct{}

func (cate *category) GetList(params model.GetCategoryListReq) model.Result[[]model.CategoryEntity] {
	list, rdCount, err := dao.Category.GetList(params)
	if err != nil {
		return model.Error[[]model.CategoryEntity](util.ErrCodeDb.ToInt(), util.ErrCodeDb.GetText())
	}

	if params.IsReturnPage == false {
		return model.Success(list)
	}

	pageCount := int(math.Ceil(float64(rdCount) / float64(params.PageSize)))
	return model.SuccessWithPage(list, params.PageIndex, params.PageSize, pageCount, rdCount)
}
