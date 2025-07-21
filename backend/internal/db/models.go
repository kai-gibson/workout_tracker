package db

import (
	"time"
)

type User struct {
	ID           int64  `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}

type Excercise struct {
	ID                  int64  `json:"id" db:"id"`
	ExcerciseTemplateID int64  `json:"excercise_template_id" db:"excercise_template_id"`
	AmountMin           uint16 `json:"amount_min" db:"amount_min"`
	AmountMax           uint16 `json:"amount_max" db:"amount_max"`
	SetMin              uint16 `json:"set_min" db:"set_min"`
	SetMax              uint16 `json:"set_max" db:"set_max"`
	SortOrder           uint16 `json:"sort_order" db:"sort_order"`
}

type Set struct {
	ID          int64     `json:"id" db:"id"`
	ExcerciseID int64     `json:"excercise_id" db:"excercise_id"`
	ProgramID   int64     `json:"program_id" db:"program_id"`
	EntryID     int64     `json:"entry_id" db:"entry_id"`
	NoteID      int64     `json:"note_id" db:"note_id"`
	Weight      uint16    `json:"weight" db:"weight"`
	Reps        uint16    `json:"reps" db:"reps"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
