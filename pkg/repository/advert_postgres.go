package repository

import (
	"github.com/amrchnk/advertisement_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"fmt"
	_"errors"
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

func (r *AdvertPostgres) GetAdvertById(id int)(models.Advert,error){
    var advert models.Advert
    query:=fmt.Sprintf("SELECT title,description,price FROM %s WHERE id=$1",advertsTable)

    err:=r.db.Get(&advert,query,id)
    return advert,err
}

func (r *AdvertPostgres) GetAllAdverts(input models.GetAdvertsFields)([]models.Advert,error){
    var adverts []models.Advert
    s:=""
    sort:=input.SortBy
    switch{
        case sort=="price":
            s+=" ORDER BY price"
        case sort=="date":
            s+=" ORDER BY created"
    }
    d:=input.Direction
    switch {
        case d=="up":
            s+=" ASC"
        default:
            s+=" DESC"
    }
    queryAdverts:=fmt.Sprintf("SELECT id,title,price FROM %s"+s,advertsTable)

    if err:=r.db.Select(&adverts,queryAdverts); err!=nil{
        return adverts,err
    }
    return adverts,nil
}