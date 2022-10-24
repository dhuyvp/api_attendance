package models

type User struct {
	Id        int    `db:"Id,omitempty" json:"Id,omitempty"`
	UserID    int    `db:"UserID,omitempty" json:"UserID,omitempty"`
	CreatedAt string `db:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	CountDay  int    `db:"CountDay,omitempty" json:"CountDay,omitempty"`
	Coin      int    `db:"Coin,omitempty" json:"Coin,omitempty"`
}
