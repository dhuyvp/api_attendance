package models

type User struct {
	Id        int    `db:"Id,omitempty" json:"id,omitempty"`
	UserID    int    `db:"UserID,omitempty" json:"user_id,omitempty"`
	CreatedAt string `db:"CreatedAt,omitempty" json:"created_at,omitempty"`
	CountDay  int    `db:"CountDay,omitempty" json:"count_day,omitempty"`
	Coin      int    `db:"Coin,omitempty" json:"coin,omitempty"`

	Day          string `json:"day,omitempty"`
	IsAttendance bool   `json:"is_attendance"`
}
