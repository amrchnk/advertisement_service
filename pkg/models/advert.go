package models

import (
	"time"
)

type Advert struct{
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
	Photos []string
	Create time.Time `json:"create"`
}

func (m *Advert) ValidateFields()string{
    if len(m.Title)>200{
        return "Ошибка: в заголовке объявления не может быть больше 200 символов"
    }

    if len(m.Description)>1000{
        return "Ошибка: в описании объявления не может быть больше 1000 символов"
    }

    if len(m.Photos>3){
        return "Ошибка: в объявлении не может быть больше 3 картинок"
    }

    if m.Title==nil || m.Price==nil || len(m.Photos)==0{
        return "Ошибка: обязательные поля не заполнены"
    }

    return nil
}