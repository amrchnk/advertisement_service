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
//     if ok:=advert.ValidateFields();ok!=nil{
//         return 0,ok
//     }
    id,err:=s.repo.CreateAdvert(advert);
    if err!=nil{
        return s.repo.CreateAdvert(advert)
    }

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

func (s *AdvertService) GetAdvertById(id int,fields []string)(map[string]interface{},error){

    advert,err:=s.repo.GetAdvertById(id)
    if err!=nil{
        return map[string]interface{}{},err
    }

    photos,err:=s.repo.GetAllPhotos(id)
    if err!=nil{
        return map[string]interface{}{},err
    }

    var res=map[string]interface{}{
        "title":advert.Title,
        "price":advert.Price,
        "photos":photos[0].Link,
    }

    if len(fields)!=0{
        for _,item:= range fields {
            if item=="photos"{
                var mas []string
                for _,photo:= range photos{
                    mas=append(mas,photo.Link)
                }
                res["photos"]=mas
            }
            if item=="description"{
                res["description"]=advert.Description
            }
        }
    }

    return res,err
}

func (s *AdvertService) GetAllAdverts(input models.GetAdvertsFields)([]map[string]interface{},error){
    var res []map[string]interface{}
    adverts,err:=s.repo.GetAllAdverts(input)
    if err!=nil{
        return res,err
    }
    for _,v:=range adverts{
        photo,err:=s.repo.GetMainPhoto(v.Id)
        if err!=nil{
            return res,err
        }
        res=append(res,map[string]interface{}{
            "title":v.Title,
            "price":v.Price,
            "photo":photo.Link,
        })
    }

    return res,err
}