package models

import (
	"time"
	"errors"
)

type Advert struct{
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price *int `json:"price"`
	Photos []string `json:"photos"`
	Create time.Time `json:"create"`
}

func (m *Advert) ValidateFields()error{
    if len(m.Title)>200{
        return errors.New("Ошибка: в заголовке объявления не может быть больше 200 символов")
    }

    if len(m.Description)>1000{
        return errors.New("Ошибка: в описании объявления не может быть больше 1000 символов")
    }

    if len(m.Photos)>3{
        return errors.New("Ошибка: в объявлении не может быть больше 3 картинок")
    }

    if m.Title=="" || m.Price==nil || len(m.Photos)==0{
        return errors.New("Ошибка: обязательные поля не заполнены")
    }

    return nil
}