package admin

import (
	"github.com/daobin/fish-blog/model"
	"github.com/daobin/fish-blog/service"
	"github.com/daobin/fish-blog/util"
	"github.com/daobin/goeasy"
)

type article struct{}

// Add godoc
// @Tags 【后台】文章管理
// @Summary 添加文章
// @Description 添加文章
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveArticleReq true "添加文章请求"
// @Success	200 {object} model.Result[model.ArticleEntity]
// @Router /admin/article [post]
func (art *article) Add(c *goeasy.Context) model.Result[any] {
	params := model.SaveArticleReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	if params.CateId == "" {
		return model.Error[any](util.CommonError.ToInt(), "分类不能为空")
	}
	if params.Title == "" {
		return model.Error[any](util.CommonError.ToInt(), "标题不能为空")
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

	return service.Article.Add(params)
}

// Update godoc
// @Tags 【后台】文章管理
// @Summary 修改文章
// @Description 修改文章
// @Accept	application/json
// @Produce json
// @Param	params body model.SaveArticleReq true "修改文章请求"
// @Success	200 {object} model.Result[model.ArticleEntity]
// @Router /admin/article [put]
func (art *article) Update(c *goeasy.Context) model.Result[any] {
	params := model.SaveArticleReq{}
	if err := c.BindJson(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	if params.ArticleId == "" {
		return model.Error[any](util.CommonError.ToInt(), "文章不能为空")
	}

	if params.State != 0 && util.CheckState(params.State) == false {
		return model.Error[any](util.CommonError.ToInt(), "状态无效")
	}

	return service.Article.Update(params)
}

// GetInfo godoc
// @Tags 【后台】文章管理
// @Summary 获取文章详情
// @Description 获取文章详情
// @Accept	application/json
// @Produce json
// @Success	200 {object} model.Result[model.ArticleEntity]
// @Router /admin/article/:id [get]
func (art *article) GetInfo(c *goeasy.Context) model.Result[any] {
	id := c.Param("id")
	if id == "" {
		return model.Error[any](util.CommonError.ToInt(), "文章不能为空")
	}

	return service.Article.GetByArticleId(id)
}

// GetList godoc
// @Tags 【后台】文章管理
// @Summary 获取文章列表
// @Description 获取文章列表
// @Accept	application/json
// @Produce json
// @Param   name query string false "名称"
// @Param   parentId query string false "父类ID"
// @Param   state query int false "状态：1（启用），2（禁用）"
// @Param   isReturnPage query bool false "是否返回分页，默认：否"
// @Param   pageIndex query int false "查询页码，默认：1"
// @Param   pageSize query int false "每页数量，默认：30"
// @Success	200 {object} model.Result[[]model.ArticleEntity]
// @Router /admin/article/list [get]
func (art *article) GetList(c *goeasy.Context) model.Result[any] {
	params := model.GetArticleListReq{}
	if err := c.BindQuery(&params); err != nil {
		return model.Error[any](util.ReqBindError.ToInt(), err.Error())
	}

	return service.Article.GetList(params)
}
