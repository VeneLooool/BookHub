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
	getReposForUser = `SELECT repo_id,
								name,
								visible,
								description,
								user_id
				FROM repos WHERE user_id = $1`
	updateUser = `UPDATE users
					SET	name = COALESCE(NULLIF($1, ''), name),
					username = COALESCE(NULLIF($2, ''), username),
					password = COALESCE(NULLIF($3, ''), password),
					user_desc = COALESCE(NULLIF($4, ''), user_desc)
					WHERE user_id = $5`
	deleteUser = `DELETE FROM users WHERE user_id = $1`
)
