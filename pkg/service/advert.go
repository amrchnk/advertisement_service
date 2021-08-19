package service

import (
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/amrchnk/advertisement_service/pkg/repository"
)

type AdvertService struct{
    repo *repository.Repository
}

func NewAdvertService(repo *repository.Repository)*AdvertService{
    return &AdvertService{repo:repo}
}

func (s *AdvertService)CreateAdvert(advert models.Advert)(int,error){
    if ok:=advert.ValidateFields();ok!=nil{
        return 0,ok
    }

    id,err:=s.repo.CreateAdvert(advert)

    for i,item:=range advert.Photos{
        var photo = models.Photo{
            Link:item,
            First:false,
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

func (s *AdvertService) GetAdvertById(id,fields []string)(models.Advert,error){
    var advert models.Advert

    if ad,err:=s.repo.GetAdvertById(id);err!=nil{
        return advert,err
    }



}