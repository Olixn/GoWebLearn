package logic

import (
	"github.com/Olixn/GoWebLearn/dao/mysql"
	"github.com/Olixn/GoWebLearn/dao/redis"
	"github.com/Olixn/GoWebLearn/models"
	"github.com/Olixn/GoWebLearn/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 生成post_id
	id := snowflake.GenID()
	p.ID = id
	// 保存到数据可
	err := mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err != nil {
		return err
	}
	return err
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

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
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
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}
		// 追加到data
		data = append(data, postData)
	}
	return
}

func GetCommunityPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
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
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}
		// 追加到data
		data = append(data, postData)
	}
	return
}

func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityPostList2(p)
	}

	if err != nil {
		zap.L().Error("GetPostListNew() failed.", zap.Error(err))
		return
	}
	return
}
