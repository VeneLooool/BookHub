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
	updateRepo = `UPDATE repos SET 
						name = COALESCE(NULLIF($1, ''), name), 
						visible = COALESCE(NULLIF($2, ''), visible),
						repo_desc = COALESCE(NULLIF($3, ''), repo_desc)
					WHERE repo_id = $4`
	deleteRepo = `DELETE FROM repos WHERE repo_id = $1`

	createBook = `INSERT INTO book (title, author, number_pages, desc, image_file_link, pdf_file_link) 
   				  VALUES ($1, $2, $3, NULLIF($4, ''), NULLIF($5, ''), $6) RETURNING book_id`
	getBook    = `SELECT book_id, title, author, number_pages, desc, image_file_link, pdf_file_link FROM books WHERE book_id = $1`
	updateBook = `UPDATE books SET
						title = COALESCE(NULLIF($1, ''), title), 
						author = COALESCE(NULLIF($2, ''), author),
						number_pages = COALESCE(NULLIF($3, ''), number_pages)
						desc = COALESCE(NULLIF($4, ''), desc), 
						image_file_link = COALESCE(NULLIF($5, ''), image_file_link),
						pdf_file_link = COALESCE(NULLIF($6, ''), pdf_file_link)
				  WHERE book_id = $7`
	deleteBook = `DELETE FROM books WHERE book_id = $1`

	attachBookToRepo = `INSERT INTO repo_books (book_id, current_page, repo_id) 
						VALUES ($1, $2, $3) RETURNING *`
	deleteBookFromRepo = `DELETE FROM repo_books WHERE book_id = &1 AND repo_id = $2`

	deleteBookFromAllRepos = `DELETE FROM repo_books WHERE book_id = &1`
)
