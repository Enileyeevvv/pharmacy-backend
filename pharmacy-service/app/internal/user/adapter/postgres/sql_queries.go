package postgres

const (
	queryCreateUser = `
		insert into users
		(login, password, status) 
		values ($1, $2, 1);
`
)
