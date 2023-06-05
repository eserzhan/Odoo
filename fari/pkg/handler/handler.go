package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/eserzhan/rest/pkg/service"
)
type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine{
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.sign_up)
		auth.POST("/sign-in", h.sign_in)
	}
	api := r.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getList)
			lists.GET("/:id", h.getListById)
			lists.DELETE("/:id", h.deleteList)
			lists.PUT("/:id", h.updateListById)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getItems)
			}
		}
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.DELETE("/:id", h.deleteItem)
			items.PUT("/:id", h.updateItem)
		}
	}

	return r
}