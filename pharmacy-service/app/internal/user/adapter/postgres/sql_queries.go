package postgres

const (
	queryCheckIfUserExists = `
		select exists(select 1
					  from users
					  where login = $1);
`

	queryCreateUser = `
		insert into users
		(login, password, status) 
		values ($1, $2, 1);
`

	queryGetPassword = `
		select password 
		from users 
		where login = $1;
`

	queryGetUser = `
		select id,
			   login,
			   password,
			   status,
			   created_at,
			   updated_at,
			   role_id
		from users
		where id = $1;
`
)
