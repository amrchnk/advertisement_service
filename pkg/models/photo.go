package models

type Photo struct {
	Id int
	Link string `db:"link"`
	First bool `db:"first"`
}
