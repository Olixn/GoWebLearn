package logic

import (
	"github.com/Olixn/GoWebLearn/dao/mysql"
	"github.com/Olixn/GoWebLearn/models"
	"github.com/Olixn/GoWebLearn/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 生成post_id
	id := snowflake.GenID()
	p.ID = id
	// 保存到数据可
	return mysql.CreatePost(p)
}

func GetPostDetailByID(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询数据并组合接口相应的数据
	post, err := mysql.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(pid) failed.", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据作者id查询作者id
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID) failed.", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	// 根据社区id查询社区详情
	communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed.", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}

	// 合并数据
	data = new(models.ApiPostDetail)
	data.AuthName = user.Username
	data.Post = post
	data.CommunityDetail = communityDetail

	// 返回
	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者id查询作者id
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed.", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区id查询社区详情
		communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed.", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		// 合并数据
		postData := &models.ApiPostDetail{
			AuthName:        user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		// 追加到data
		data = append(data, postData)
	}
	return
}
