package admin

import (
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/service"
	"github.com/daobin/fish-blog/util"
	"github.com/daobin/goeasy"
)

type category struct{}

// Add godoc
// @Tags 【后台】文章分类管理
// @Summary 添加文章分类
// @Description 添加文章分类
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveCategoryReq true "添加文章分类请求"
// @Success	200 {object} model.Result[model.CategoryResp]
// @Router /admin/category [post]
func (cate *category) Add(c *goeasy.Context) model.Result[any] {
	params := model.SaveCategoryReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ErrCodeReqBind.ToInt(), err.Error())
	}

	return model.Success[any](nil)
}

// Update godoc
// @Tags 【后台】文章分类管理
// @Summary 更新文章分类
// @Description 更新文章分类
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveCategoryReq true "更新文章分类请求"
// @Success	200 {object} model.Result[model.CategoryResp]
// @Router /admin/category [put]
func (cate *category) Update(c *goeasy.Context) model.Result[any] {
	params := model.SaveCategoryReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ErrCodeReqBind.ToInt(), err.Error())
	}

	return model.Success[any](nil)
}

// Delete godoc
// @Tags 【后台】文章分类管理
// @Summary 添加文章分类
// @Description 添加文章分类
// @Accept	application/json
// @Produce json
// @Param	params body model.ObjectIdReq true "删除文章分类请求"
// @Success	200 {object} model.Result[any]
// @Router /admin/category [delete]
func (cate *category) Delete(c *goeasy.Context) model.Result[any] {
	params := model.ObjectIdReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ErrCodeReqBind.ToInt(), err.Error())
	}

	return model.Success[any](nil)
}

// GetInfo godoc
// @Tags 【后台】文章分类管理
// @Summary 添加文章分类
// @Description 添加文章分类
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveCategoryReq true "添加文章分类请求"
// @Success	200 {object} model.Result[model.CategoryResp]
// @Router /admin/category [get]
func (cate *category) GetInfo(c *goeasy.Context) model.Result[any] {

	return model.Success[any](nil)
}

// GetList godoc
// @Tags 【后台】文章分类管理
// @Summary 挂载网卡
// @Description 挂载网卡
// @Accept	application/json
// @Produce json
// @Param   name query string false "名称"
// @Param   state query int false "状态：1（启用），2（禁用）"
// @Param   isReturnPage query bool false "是否返回分页，默认：否"
// @Param   pageIndex query int false "查询页码，默认：1"
// @Param   pageSize query int false "每页数量，默认：30"
// @Success	200 {object} model.Result[[]model.CategoryResp]
// @Router /admin/category/list [get]
func (cate *category) GetList(c *goeasy.Context) model.Result[any] {
	params := model.GetCategoryListReq{}
	if err := c.BindQuery(&params); err != nil {
		return model.Error[any](util.ErrCodeReqBind.ToInt(), err.Error())
	}

	return service.Category.GetList(params).ToAny()
}
