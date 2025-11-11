package posts

import "time"

type (
	UserActivityRequest struct {
		IsLiked bool `json:"isLiked"`
	}
)

type UserActivityModel struct {
	ID        int64     `db:"id"`
	PostID    int64     `db:"id"`
	UserID    int64     `db:"id"`
	IsLiked   bool      `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}
