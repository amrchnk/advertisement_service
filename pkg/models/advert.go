package models

import (
    "errors"
    "strings"
//     "strconv"
)

type Advert struct{
	Id int `json:"-" db:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
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

    if m.Title=="" || m.Price==0 || len(m.Photos)==0{
        return errors.New("Invalid data")
    }

    return nil
}

type GetAdvertsFields struct{
    Page int `json:"page"`
    SortBy string `json:"sortBy"`
    Direction string `json:"direction"`
}

func (af *GetAdvertsFields) ValidateInput()(string,bool){
    if !(strings.ToLower(af.SortBy)=="date"||strings.ToLower(af.SortBy)=="price"){
        return "sort incorrect",false
    }

    if !(strings.ToLower(af.Direction)!="up"||strings.ToLower(af.Direction)!="down"){
        return "direct incorrect",false
    }
    return "ok",true
}