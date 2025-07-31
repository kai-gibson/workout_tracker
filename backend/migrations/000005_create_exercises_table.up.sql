CREATE TABLE tbl_exercises (
    id BIGSERIAL PRIMARY KEY,
    exercise_template_id BIGINT REFERENCES tbl_exercise_templates(id),
    program_entry_id BIGINT REFERENCES tbl_program_entries(id),
    name VARCHAR(60) NOT NULL,
    amount_min SMALLINT NOT NULL,
    amount_max SMALLINT NOT NULL,
    set_min SMALLINT NOT NULL,
    set_max SMALLINT NOT NULL,
    sort_order SMALLINT NOT NULL
)
