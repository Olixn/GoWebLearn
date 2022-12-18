package controller

import (
	"github.com/Olixn/GoWebLearn/logic"
	"github.com/Olixn/GoWebLearn/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVote)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed.", zap.Any("err", err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			errData := removeTopStruct(errs.Translate(trans))
			ResponseErrorWithMsg(c, CodeInvalidParam, errData)
			return
		}
	}

	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.UserID = userID

	err = logic.VoteForPost(p)
	if err != nil {
		zap.L().Error("logic.VoteForPost(p) failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
	return
}
