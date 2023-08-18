package router

import (
	"Moonlay/handler"
	"Moonlay/pkg/middleware"
	"Moonlay/pkg/mysql"
	"Moonlay/repository"

	"github.com/labstack/echo/v4"
)

func SubList(r *echo.Group) {
	subListRepository := repository.SubListRepository(mysql.DB)
	h := handler.HandlerSubList(subListRepository)

	r.GET("/sub_list", h.GetSubList)
	r.GET("/sub_list/:id", h.GetSubByID)
	r.POST("/sub_list", middleware.UploadFile(h.CreateSubList))
	r.PATCH("/sub_list/:id", middleware.UploadFile(h.UpdateSubList))
	r.DELETE("/sub_list/:id", h.DeleteSubList)
}
