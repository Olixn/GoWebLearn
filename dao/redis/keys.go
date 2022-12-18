package redis

// redis key

// redis key 注意使用命名空间的方式，方便业务去拆分

const (
	Prefix             = "bluebell:"
	KeyPostTimeZSet    = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;帖子及投票分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型,参数是帖子post_id
	KeyCommunitySetPF  = "community:"
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
