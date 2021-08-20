package models

import (
    "errors"
)

type Advert struct{
	Id int `json:"-" db:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price *int `json:"price"`
	Photos []string `json:"photos"`
}

func (m *Advert) ValidateFields()error{
    if len(m.Title)>200{
        return errors.New("Invalid data")
    }

    if len(m.Description)>1000{
        return errors.New("Invalid data")
    }

    if len(m.Photos)>3{
        return errors.New("Invalid data")
    }

    if m.Title=="" || m.Price==nil || len(m.Photos)==0{
        return errors.New("Invalid data")
    }

    return nil
}