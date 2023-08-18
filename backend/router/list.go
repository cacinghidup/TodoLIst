package router

import (
	"Moonlay/handler"
	"Moonlay/pkg/middleware"
	"Moonlay/pkg/mysql"
	"Moonlay/repository"

	"github.com/labstack/echo/v4"
)

func ListRouter(e *echo.Group) {
	listRepository := repository.ListRepository(mysql.DB)
	h := handler.HandlerList(listRepository)

	e.GET("/list", h.GetList)
	e.GET("/list/:id", h.GetByID)
	e.POST("/list", middleware.UploadFile(h.CreateList))
	e.PATCH("/list/:id", middleware.UploadFile(h.UpdateList))
	e.DELETE("/list/:id", h.DeleteList)
}
