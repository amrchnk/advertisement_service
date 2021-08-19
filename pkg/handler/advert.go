package handler

import (
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
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

