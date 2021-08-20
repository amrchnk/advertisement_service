package repository

import (
	"github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/jmoiron/sqlx"
	"fmt"
)

type PhotoPostgres struct{
	db *sqlx.DB
}

func NewPhotoPostgres(db *sqlx.DB) *PhotoPostgres{
	return &PhotoPostgres{db:db}
}

func (r *PhotoPostgres) CreatePhoto(photo models.Photo,adv_id int)(int,error){
	//Transaction start
	tx,err:=r.db.Begin()
	if err!=nil{
		return 0,err
	}

	var id int
	createItemQuery:=fmt.Sprintf("INSERT INTO %s (link,first,advert_id) VALUES ($1,$2,$3) RETURNING id",photosTable)

	row:=tx.QueryRow(createItemQuery,photo.Link,photo.First,adv_id)
	err=row.Scan(&id)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}

	return id,tx.Commit()
}

func (r *PhotoPostgres) GetMainPhoto(advertId int)(models.Photo,error){
    var photo models.Photo
    query:=fmt.Sprintf("SELECT link FROM %s WHERE advert_id=$1 AND first=true",photosTable)
    err:=r.db.Get(&photo,query,advertId)
    return photo,err
}

func (r *PhotoPostgres) GetAllPhotos(advertId int)([]models.Photo,error){
    var photos []models.Photo
    query:=fmt.Sprintf("SELECT link FROM %s WHERE advert_id=$1",photosTable)
    err:=r.db.Select(&photos,query,advertId)

    return photos,err
}



