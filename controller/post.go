package controller

import (
	"strconv"

	"github.com/Olixn/GoWebLearn/logic"
	"github.com/Olixn/GoWebLearn/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数并校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error.", zap.Any("err", err))
		zap.L().Error("create post with invalid param.")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从 c 拿到author_id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")

	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param.", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetail(pid) failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// PostListHandler 帖子分页列表
func PostListHandler(c *gin.Context) {

	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// PostListHandler2 帖子分页列表
func PostListHandler2(c *gin.Context) {

	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery(p); failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostListNew(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

//// CommunityPostListHandler 根据社区去查询帖子列表
//func CommunityPostListHandler(c *gin.Context) {
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("c.ShouldBindQuery(p); failed", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	data, err := logic.GetCommunityPostList2(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//
//	ResponseSuccess(c, data)
//}
