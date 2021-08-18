package service

import (
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/amrchnk/advertisement_service/pkg/repository"
    "errors"
)

type AdvertService struct{
    repo repository.Repository
}

func NewAdvertService(repo repository.Advert)*AdvertService{
    return &AdvertService{repo:repo}
}

func (s *AdvertService)CreateAdvert(advert models.Advert)(int,error){

    if ok:=advert.ValidateFields();ok!=nil{
        return 0,errors.New(ok)
    }

    id,err:=s.repo.CreateAdvert(advert)

    for i,item:=range advert.Photos{
        var photo models.Photo{
            Link:item,
            First:false
        }

        if i==0{
            photo.First=true
        }

        if _,err=s.repo.CreatePhoto(photo,id);err!=nil{
            return 0,err
        }
    }
    return id,err
}