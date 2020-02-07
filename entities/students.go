package entities

type Student struct {
	ID         int64  `json:"id"`
	RollID     int64  `json:"roll_id"`
	Name       string `json:"std_name"`
	FatherName string `json:"father_name"`
	MotherName string `json:"mother_name"`
	Degree     string `json:"degree"`
	Year       int64  `json:"year"`

	Board       string `json:"board"`
	Institute   string `json:"institute"`
	Bangla      int64  `json:"bangla"`
	GeneralMath int64  `json:"general_math"`
	English     int64  `json:"eng"`
	History     int64  `json:"history"`
	CreatedAt   string `json:"created_at"`
}
