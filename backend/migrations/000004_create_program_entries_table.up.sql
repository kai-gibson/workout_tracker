CREATE TABLE tbl_program_entries (
    id BIGSERIAL PRIMARY KEY,
    program_id BIGINT REFERENCES tbl_programs(id),
    sort_order SMALLINT NOT NULL
)
