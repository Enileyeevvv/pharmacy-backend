package postgres

type User struct {
	ID        int    `db:"id"`
	Login     string `db:"login"`
	Password  string `db:"password"`
	Status    int    `db:"status"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
	RoleID    int    `db:"role_id"`
}
