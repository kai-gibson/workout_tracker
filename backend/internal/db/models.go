package db

import (
	"time"
)

type User struct {
	ID             int64  `json:"id" db:"id"`
	Username       string `json:"username" db:"username"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
}

type ExcerciseTemplate struct {
	ID           int64   `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	MuscleGroups []int32 `json:"muscle_groups" db:"muscle_groups"`
	ImageURL     string  `json:"image_url" db:"image_url"`
}

type Program struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserID int64  `json:"user_id" db:"user_id"`
}

type ProgramEntry struct {
	ID        int64 `json:"id" db:"id"`
	ProgramID int64 `json:"program_id" db:"program_id"`
	SortOrder int16 `json:"sort_order" db:"sort_order"`
}

type Excercise struct {
	ID                  int64 `json:"id" db:"id"`
	ExcerciseTemplateID int64 `json:"excercise_template_id" db:"excercise_template_id"`
	ProgramEntryID      int64 `json:"program_entry_id" db:"program_entry_id"`
	AmountMin           int16 `json:"amount_min" db:"amount_min"`
	AmountMax           int16 `json:"amount_max" db:"amount_max"`
	SetMin              int16 `json:"set_min" db:"set_min"`
	SetMax              int16 `json:"set_max" db:"set_max"`
	SortOrder           int16 `json:"sort_order" db:"sort_order"`
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
