package service

import (
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/amrchnk/advertisement_service/pkg/repository"
)

type Advert interface{
    CreateAdvert(advert models.Advert)(int,error)
    GetAdvertById(id,fields []string)(advert models.Advert,error)
}

type Service struct{
    Advert
}

func NewService(repos *repository.Repository) *Service{
    return &Service{
        Advert: NewAdvertService(repos),
    }
}