CREATE TABLE tbl_users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    hashed_password VARCHAR(200) NOT NULL
)
