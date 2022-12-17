package mysql

import (
	"github.com/Olixn/GoWebLearn/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO post (post_id, title, content, author_id, community_id) VALUE (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)

	return
}

func GetPostDetailByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time FROM post WHERE post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time FROM post ORDER BY id LIMIT ?,?`

	posts = make([]*models.Post, 0, size) // 不是这样make([]*models.Post, 2)

	err = db.Select(&posts, sqlStr, page-1, size)
	return

}
