CREATE TABLE tbl_programs (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    user_id BIGINT REFERENCES tbl_users(id)
)
