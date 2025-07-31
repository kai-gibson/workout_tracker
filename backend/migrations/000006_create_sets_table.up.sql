CREATE TABLE tbl_sets (
    id           BIGSERIAL    PRIMARY KEY,
    excercise_id BIGINT       REFERENCES tbl_exercises(id),
    program_id   BIGINT       REFERENCES tbl_programs(id),
    entry_id     BIGINT       REFERENCES tbl_program_entries(id),
    weight       SMALLINT     NOT NULL,
    reps         SMALLINT     NOT NULL,
    created_at   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);
