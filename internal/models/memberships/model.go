package memberships

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type (
	UserModel struct {
		ID        int64  `db:"id"`
		Email     string `db:"email"`
		Username  string `db:"username"`
		Password  string `db:"password"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:" updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:" updated_by"`
	}
)
