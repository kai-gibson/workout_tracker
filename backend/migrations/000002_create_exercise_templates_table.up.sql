CREATE TABLE tbl_exercise_templates (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    muscle_groups INTEGER[] NOT NULL,
    image_url VARCHAR(255)
)
