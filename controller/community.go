package controller

import (
	"strconv"

	"github.com/Olixn/GoWebLearn/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 社区相关路由处理

// CommunityHandler 社区分类列表
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name),以列表的形式返回

	if data, err := logic.GetCommunityList(); err != nil {
		zap.L().Error("logic.GetCommunityList() failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		ResponseSuccess(c, data)
	}
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	communityID := c.Param("id")
	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	if data, err := logic.GetCommunityDetail(id); err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		ResponseSuccess(c, data)
	}
}
