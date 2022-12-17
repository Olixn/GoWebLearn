package mysql

import (
	"database/sql"

	"github.com/Olixn/GoWebLearn/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `SELECT community_id,community_name FROM community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db.")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `SELECT community_id,community_name,introduction,create_time FROM community WHERE community_id = ?`
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community_id in db.", zap.Error(err))
			err = ErrorInvalidID
		}
	}
	return
}
