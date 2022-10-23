package postgres

const (
	createUser = `INSERT INTO users (name, username, password, user_desc) 
					VALUES (NULLIF($1, ''), $2, $3, NULLIF($4, '')) 
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

	createRepo = `INSERT INTO repos (name, visible, repo_desc, user_id)
					VALUES (NULLIF($1, ''), $2, NULLIF($3, ''), $4) 
					RETURNING *`
	getRepo = `SELECT 	repo_id,
						name,
						visible,
						repo_desc,
						user_id,
				FROM repos WHERE repo_id = $1`
	getBooksForRepo = `SELECT 
						books.book_id, 
						books.title, 
						books.author, 
						book.number_pages, 
						books.desc, 
						repo_books.current_page,
					   FROM books INNER JOIN (SELECT * FROM repo_books WHERE repo_id = $1) on books.book_id = repo_books.id`
	updateRepo = `UPDATE repos
					SET name = COALESCE(NULLIF($1, ''), name), 
						visible = COALESCE(NULLIF($2, ''), visible),
						repo_desc = COALESCE(NULLIF($3, ''), repo_desc)
					WHERE repo_id = $4`
	deleteRepo = `DELETE FROM repos WHERE repo_id = $1`
)
