package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type AdvertPostgres struct{
	db *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres{
	return &AdvertPostgres{db:db}
}

func (r *AdvertPostgres) CreateAdvert(advert models.Advert)(int,error){
	//Transaction start
	tx,err:=r.db.Begin()
	if err!=nil{
		return 0,err
	}

	var id int
	createItemQuery:=fmt.Sprintf("INSERT INTO %s (title,description,price) VALUES ($1,$2,$3) RETURNING id",advertsTable)

	row:=tx.QueryRow(createItemQuery,advert.Title,advert.Description,advert.Price)
	err=row.Scan(&id)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}

	return id,tx.Commit()
}