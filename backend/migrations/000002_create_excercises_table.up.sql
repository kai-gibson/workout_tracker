CREATE TABLE tbl_excercises (
    id BIGSERIAL PRIMARY KEY,
    excercise_template_id BIGINT,
    amt_min TINYINT,
    amt_max TINYINT,
    set_min TINYINT,
    set_max TINYINT,
    sort_order TINYINT
)
