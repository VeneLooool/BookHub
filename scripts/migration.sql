create table if not exists users(
                                    user_id serial not null primary key,
                                    name text,
                                    username text,
                                    password text,
                                    user_desc text
);

create table if not exists books (
                                     book_id serial not null primary key,
                                     title text,
                                     author text,
                                     number_pages int,
                                     description text,
                                     image_file_link text,
                                     pdf_file_link text
);

create table if not exists repos (
                                     repo_id serial not null primary key,
                                     name text,
                                     visible boolean,
                                     repo_desc text,
                                     user_id  int references users (user_id)
    );

create table if not exists repo_books (
    book_id int references books (book_id),
    current_page int,
    repo_id int references repos (repo_id)
    );