package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/amrchnk/advertisement_service/pkg/models"
)

type Advert interface{
	CreateAdvert(advert models.Advert)(int,error)
}

type Photo interface{
	CreatePhoto(photo models.Photo,adv_id int)(int,error)
}

type Repository struct {
	Advert
	Photo
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Advert: NewAdvertPostgres(db),
		Photo: NewPhotoPostgres(db),
	}
}

