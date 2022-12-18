package logic

import (
	"strconv"

	"go.uber.org/zap"

	"github.com/Olixn/GoWebLearn/dao/redis"
	"github.com/Olixn/GoWebLearn/models"
)

// VoteForPost PostForVote 为帖子投票
func VoteForPost(p *models.ParamVote) error {
	zap.L().Debug("VoteForPost.", zap.Int64("userID", p.UserID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(p.UserID)), p.PostID, float64(p.Direction))
}
