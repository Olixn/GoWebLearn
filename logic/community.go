package logic

import (
	"github.com/Olixn/GoWebLearn/dao/mysql"
	"github.com/Olixn/GoWebLearn/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库，查找所有的community 并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetail 根据社区ＩＤ查询社区详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
