package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// 投票功能

// 投一票加432分 86400/200 -> 需要200张赞成才可以为你的帖子续一天

// 投票几种情况：
// direction=1 时:
//		1. 之前没投过票，现在投赞成 -> 更新分数和投票纪录 1 +432
// 		2. 之前投反对，现在改赞成 -> 更新分数和投票纪录   2 +432*2
// direction=0 时:
//		1. 之前投赞成，现在改取消 -> 更新分数和投票纪录  1 -432
// 		2. 之前投反对，现在改取消 -> 更新分数和投票纪录  1 +432
// direction=-1 时:
//		1. 之前没投过票，现在投反对 -> 更新分数和投票纪录 1 -432
// 		2. 之前投赞成，现在改反对 -> 更新分数和投票纪录   2 -432*2

// 投票限制：
// 每个帖子自发表之日一个星期内允许投票，否则允许
// 截止投票后，将数据存入MySQL 中
// 截止投票后，删除KeyPostVotedZSetPF

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {
	// 事务操作
	pipeline := rdb.TxPipeline()

	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票限制
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2.更新分数
	// 先查当前用户之前给当前帖子的投票纪录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()

	if value == ov {
		return ErrVoteRepeated
	}

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值

	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 记录用户为该帖子投票
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成还是反对
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
