package admin

import (
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/service"
	"github.com/daobin/fish-blog/util"
	"github.com/daobin/goeasy"
)

type category struct{}

// Add godoc
// @Tags 【后台】分类管理
// @Summary 添加分类
// @Description 添加分类
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveCategoryReq true "添加分类请求"
// @Success	200 {object} model.Result[model.CategoryEntity]
// @Router /admin/category [post]
func (cate *category) Add(c *goeasy.Context) model.Result[any] {
	params := model.SaveCategoryReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	if params.Name == "" {
		return model.Error[any](util.CommonError.ToInt(), "名称不能为空")
	}
	if params.State == 0 {
		params.State = util.StateEnable
	}
	if util.CheckState(params.State) == false {
		return model.Error[any](util.CommonError.ToInt(), "状态无效")
	}
	if params.Sort <= 0 {
		params.Sort = util.SortDefault
	}

	return service.Category.Add(params)
}

// Update godoc
// @Tags 【后台】分类管理
// @Summary 修改分类
// @Description 修改分类
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveCategoryReq true "修改分类请求"
// @Success	200 {object} model.Result[model.CategoryEntity]
// @Router /admin/category [put]
func (cate *category) Update(c *goeasy.Context) model.Result[any] {
	params := model.SaveCategoryReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	if params.CateId == "" {
		return model.Error[any](util.CommonError.ToInt(), "分类不能为空")
	}

	if params.State != 0 && util.CheckState(params.State) == false {
		return model.Error[any](util.CommonError.ToInt(), "状态无效")
	}

	return service.Category.Update(params)
}

// Delete godoc
// @Tags 【后台】分类管理
// @Summary 删除分类
// @Description 删除分类
// @Accept	application/json
// @Produce json
// @Param	params body model.ObjectIdReq true "删除分类请求"
// @Success	200 {object} model.Result[any]
// @Router /admin/category [delete]
func (cate *category) Delete(c *goeasy.Context) model.Result[any] {
	params := model.ObjectIdReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	if params.Id == "" {
		return model.Error[any](util.CommonError.ToInt(), "分类不能为空")
	}

	return service.Category.Delete(params.Id)
}

// GetInfo godoc
// @Tags 【后台】分类管理
// @Summary 获取分类详情
// @Description 获取分类详情
// @Accept	application/json
// @Produce json
// @Param   id query string true "分类ID"
// @Success	200 {object} model.Result[model.CategoryEntity]
// @Router /admin/category [get]
func (cate *category) GetInfo(c *goeasy.Context) model.Result[any] {
	id := c.Query("id")
	if id == "" {
		return model.Error[any](util.CommonError.ToInt(), "分类不能为空")
	}

	return service.Category.GetByCateId(id)
}

// GetList godoc
// @Tags 【后台】分类管理
// @Summary 获取分类列表
// @Description 获取分类列表
// @Accept	application/json
// @Produce json
// @Param   name query string false "名称"
// @Param   parentId query string false "父类ID"
// @Param   state query int false "状态：1（启用），2（禁用）"
// @Param   isReturnPage query bool false "是否返回分页，默认：否"
// @Param   pageIndex query int false "查询页码，默认：1"
// @Param   pageSize query int false "每页数量，默认：30"
// @Success	200 {object} model.Result[[]model.CategoryEntity]
// @Router /admin/category/list [get]
func (cate *category) GetList(c *goeasy.Context) model.Result[any] {
	params := model.GetCategoryListReq{}
	if err := c.BindQuery(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	return service.Category.GetList(params)
}

// GetTree godoc
// @Tags 【后台】分类管理
// @Summary 获取分类树
// @Description 获取分类树
// @Accept	application/json
// @Produce json
// @Param   parentId query string false "父类ID"
// @Success	200 {object} model.Result[model.CategoryTreeResp]
// @Router /admin/category/tree [get]
func (cate *category) GetTree(c *goeasy.Context) model.Result[any] {
	return service.Category.GetTree(c.Query("parentId"))
}
