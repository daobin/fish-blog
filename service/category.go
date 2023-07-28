package service

import (
	"github.com/daobin/fish-blog/dao"
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/util"
	"gopkg.in/mgo.v2/bson"
	"math"
	"time"
)

type category struct{}

func (cate *category) IsExist(info *model.CategoryEntity) bool {
	return info != nil && info.State != util.StateDelete
}

func (cate *category) Add(params model.SaveCategoryReq) model.Result[any] {
	// 校验分类名称是否存在
	nameInfo, err := dao.Category.GetByCateName(params.Name, "")
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if cate.IsExist(nameInfo) {
		return model.Error[any](util.CommonError.ToInt(), "分类名称已存在")
	}

	// 校验分类父级是否存在
	if params.ParentId != "" {
		parentInfo, err := dao.Category.GetByCateId(params.ParentId)
		if err != nil {
			return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
		}
		if cate.IsExist(parentInfo) == false {
			return model.Error[any](util.CommonError.ToInt(), "分类父级不存在")
		}
	}

	entity := &model.CategoryEntity{}
	util.Struct2Struct(params, entity)

	if nameInfo != nil && nameInfo.State == util.StateDelete {
		// 恢复逻辑删除的数据
		entity.Id = nameInfo.Id
		entity.CateId = nameInfo.CateId
	} else {
		entity.Id = bson.NewObjectId()
		entity.CateId = util.GenRandomString(8)
	}
	entity.CreatedAt = time.Now().Unix()
	entity.UpdatedAt = entity.CreatedAt

	err = dao.Category.Save(entity)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	return model.Success[any](entity)
}

func (cate *category) Update(params model.SaveCategoryReq) model.Result[any] {
	// 校验分类是否存在
	info, err := dao.Category.GetByCateId(params.CateId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if cate.IsExist(info) == false {
		return model.Error[any](util.CommonError.ToInt(), "分类不存在")
	}

	// 校验分类父级是否存在
	if params.ParentId != "" {
		parentInfo, err := dao.Category.GetByCateId(params.ParentId)
		if err != nil {
			return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
		}
		if cate.IsExist(parentInfo) == false {
			return model.Error[any](util.CommonError.ToInt(), "分类父级不存在")
		}

		info.ParentId = params.ParentId
	}

	// 校验分类名称是否存在
	if params.Name != "" && params.Name != info.Name {
		nameInfo, err := dao.Category.GetByCateName(params.Name, info.CateId)
		if err != nil {
			return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
		}
		if cate.IsExist(nameInfo) {
			return model.Error[any](util.CommonError.ToInt(), "分类名称已存在")
		}

		info.Name = params.Name
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

	err = dao.Category.Save(info)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	return model.Success[any](info)
}

func (cate *category) Delete(cateId string) model.Result[any] {
	// 校验分类名称是否存在
	info, err := dao.Category.GetByCateId(cateId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if cate.IsExist(info) == false {
		return model.Success[any](nil)
	}

	info.State = util.StateDelete
	info.UpdatedAt = time.Now().Unix()

	err = dao.Category.Save(info)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	return model.Success[any](nil)
}

func (cate *category) GetByCateId(cateId string) model.Result[any] {
	info, err := dao.Category.GetByCateId(cateId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if cate.IsExist(info) == false {
		return model.Error[any](util.CommonError.ToInt(), "分类不存在")
	}

	return model.Success[any](info)
}

func (cate *category) GetList(params model.GetCategoryListReq) model.Result[any] {
	list, rdCount, err := dao.Category.GetList(&params)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	if params.IsReturnPage == false {
		return model.Success[any](list)
	}

	pageCount := int(math.Ceil(float64(rdCount) / float64(params.PageSize)))
	return model.SuccessWithPage[any](list, params.PageIndex, params.PageSize, pageCount, rdCount)
}

func (cate *category) GetTree(parentId string) model.Result[any] {
	parentInfo, err := dao.Category.GetByCateId(parentId)
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	childrenMap, err := cate.getChildrenMap()
	if err != nil {
		return model.Error[any](util.DbAbnormal.ToInt(), util.DbAbnormal.GetText())
	}

	cateTree := model.CategoryTreeResp{
		Children: []model.CategoryTreeResp{},
	}
	util.Struct2Struct(*parentInfo, &cateTree.CategoryEntity)

	cate.childTreeProcess(&cateTree, childrenMap)

	return model.Success[any](cateTree)
}

func (cate *category) getChildrenIds(parentId string) []string {
	childrenMap, _ := cate.getChildrenMap()
	if len(childrenMap) == 0 || len(childrenMap[parentId]) == 0 {
		return []string{}
	}

	cateIds := make([]string, 0, len(childrenMap[parentId]))
	if parentId != "" {
		cateIds = append(cateIds, parentId)
	}

add:
	for _, child := range childrenMap[parentId] {
		cateIds = append(cateIds, child.CateId)
		if len(childrenMap[child.CateId]) > 0 {
			parentId = child.CateId
			goto add
		}
	}

	return cateIds
}

func (cate *category) getChildrenMap() (map[string][]model.CategoryEntity, error) {
	list, rdCount, err := dao.Category.GetList(&model.GetCategoryListReq{})
	if err != nil {
		return nil, util.DbAbnormal.NewError()
	}
	if rdCount == 0 {
		return nil, nil
	}

	childrenMap := make(map[string][]model.CategoryEntity)
	for _, info := range list {
		if _, ok := childrenMap[info.ParentId]; !ok {
			childrenMap[info.ParentId] = make([]model.CategoryEntity, 0)
		}
		childrenMap[info.ParentId] = append(childrenMap[info.ParentId], info)
	}

	return childrenMap, nil
}

func (cate *category) childTreeProcess(cateTree *model.CategoryTreeResp, childrenMap map[string][]model.CategoryEntity) {
	if cateTree == nil || len(childrenMap[cateTree.CateId]) == 0 {
		return
	}

	children := make([]model.CategoryTreeResp, 0)
	for _, info := range childrenMap[cateTree.CateId] {
		tree := model.CategoryTreeResp{
			Children: []model.CategoryTreeResp{},
		}
		util.Struct2Struct(info, &tree.CategoryEntity)

		if len(childrenMap[tree.CateId]) > 0 {
			cate.childTreeProcess(&tree, childrenMap)
		}

		children = append(children, tree)
	}

	cateTree.Children = children
	delete(childrenMap, cateTree.CateId)
}
