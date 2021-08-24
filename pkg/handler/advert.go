package handler

import (
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
//     "fmt"
    "math"
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

    if err:=input.ValidateFields();err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
                "id":-1,
                "status":http.StatusBadRequest,
        })
        return
    }

    id,err:=h.services.Advert.CreateAdvert(input)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
                "id":-1,
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
    if len(a.Fields)==0{
        return true
    }

    for i:=0;i<len(a.Fields);i++{
        if a.Fields[i]=="description"||a.Fields[i]=="photos"{
            continue
        } else{
            return false
        }
    }
    return true
}



func (h *Handler) getAdvertById(c *gin.Context){
    var input AdvertFields
    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
        })
        return
    }

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
    res,err:=h.services.Advert.GetAdvertById(id,input.Fields)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "id":-1,
            "status":http.StatusInternalServerError,
        })
        return
    }

    c.JSON(http.StatusOK,res)
}

func (h *Handler) getAllAdverts(c *gin.Context){
    const pagination=10
    var input models.GetAdvertsFields

    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
            "err":"body is not valid",
        })
        return
    }
    if str,err:=input.ValidateInput();!err{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
            "err":str,
        })
        return
    }
    res,err:=h.services.GetAllAdverts(input)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusInternalServerError,
        })
        return
     }
    if (input.Page==0||int(math.Ceil(float64(len(res))/float64(pagination)))<input.Page){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
        })
        return
    }
    l:=(input.Page-1)*pagination
    r:=input.Page*pagination
    if(len(res)<r){
        res=res[l:]
    }else{
        res=res[l:r]
    }

    c.JSON(http.StatusOK,res)
}
