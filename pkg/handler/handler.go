package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/amrchnk/advertisement_service/pkg/service"
)

type Handler struct{
    services *service.Service
}

func NewHandler(services *service.Service) *Handler{
    return &Handler{services:services}
}

func (h *Handler) InitRoutes() *gin.Engine{
    router:=gin.New()

    api:=router.Group("/adverts")
    {
        api.POST("/",h.createAdvert)
        api.GET("/:id",h.getAdvertById)
    }

    return router
}