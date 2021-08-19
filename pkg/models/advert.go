package models

import (
	"time"
)

type Advert struct{
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price *int `json:"price"`
	Photos []string `json:"photos"`
}

func (m *Advert) ValidateFields()error{
    if len(m.Title)>200{
        return error
    }

    if len(m.Description)>1000{
        return error
    }

    if len(m.Photos)>3{
        return error
    }

    if m.Title=="" || m.Price==nil || len(m.Photos)==0{
        return error
    }

    return nil
}