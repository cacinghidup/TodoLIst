package router

import "github.com/labstack/echo/v4"

func Routes(r *echo.Group) {
	ListRouter(r)
	SubList(r)
}
