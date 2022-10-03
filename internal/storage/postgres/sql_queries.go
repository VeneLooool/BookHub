package postgres

const (
	createUser = `INSERT INTO users (user_id, name, username, password, user_desc) 
					VALUES ($1, NULLIF($2, ''), $3, $4, NULLIF($5, '')) 
					RETURNING *`
	getUser = `SELECT user_id,
						name,
						username,
						password,
						user_desc
				FROM users WHERE user_id = $1`
	getReposForUser = ``
)
