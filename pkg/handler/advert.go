package handler

import (
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func (h *Handler) createAdvert(c *gin.Context){
    var input models.Advert
    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
                "id":-1,
                "status":http.StatusBadRequest,
        })
        return
    }

    id,err:=h.services.Advert.CreateAdvert(input)
    if err!=nil{
            c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
                    "id":id,
                    "status":http.StatusInternalServerError,
            })
        return
    }

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":id,
        "status":http.StatusOK,
    })
}

type AdvertFields struct{
    Fields []string `json:"fields"`
}

func (a *AdvertFields) ValidateInput()bool{
    for _,item:=range a{
        if item!="description"||item!="photos"{
            return false
        }
    }
    return true
}

func (h *Handler) getAdvertById(c *gin.Context){
    var input AdvertFields

    c.BindJSON(&input)

    if !input.ValidateInput(){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
        })
        return
    }

    id,err:=strconv.Atoi(c.Param("id"))
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
        })
        return
    }

    res,err:=h.services.Advert.GetAdvertById(id,input)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusInternalServerError,
        })
        return
    }

    c.JSON(http.StatusOK,res)
}

func (h *Handler) getAllAdverts(c *gin.Context){

}
