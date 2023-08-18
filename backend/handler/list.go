package handler

import (
	listdto "Moonlay/dto/list"
	dto "Moonlay/dto/result"
	"Moonlay/models"
	"Moonlay/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerList struct {
	HandlerRepository repository.List
}

func HandlerList(HandlerRepository repository.List) *handlerList {
	return &handlerList{HandlerRepository}
}

// func (h *handlerList) GetList(c echo.Context) error {
// 	list, err := h.HandlerRepository.FindAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: list})
// }

// func (h *handlerList) GetList(c echo.Context) error {
// 	page, err := strconv.Atoi(c.QueryParam("page"))
// 	if err != nil {
// 		page = 1
// 	}

// 	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
// 	if err != nil {
// 		pageSize = 10
// 	}

// 	totalRecords, _ := h.HandlerRepository.GetTotalRecords()

// 	lists, err := h.HandlerRepository.FindAll(page, pageSize, totalRecords)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: lists})
// }

func (h *handlerList) GetList(c echo.Context) error {
	title := c.QueryParam("title")
	deskripsi := c.QueryParam("deskripsi")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		pageSize = 5
	}

	totalRecords, _ := h.HandlerRepository.GetTotalRecords()

	lists, err := h.HandlerRepository.FindAll(page, totalRecords, pageSize, title, deskripsi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: lists})
}

func (h *handlerList) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	list, err := h.HandlerRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: list})
}

func (h *handlerList) CreateList(c echo.Context) error {
	file, nf := c.Get("fileUpload").(string)
	if !nf {
		file = ""
	}

	request := listdto.CreateList{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
		Upload:    file,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}

	list := models.List{
		Title:     request.Title,
		Deskripsi: request.Deskripsi,
		Upload:    request.Upload,
	}

	data, err := h.HandlerRepository.CreateList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}

func (h *handlerList) UpdateList(c echo.Context) error {
	file, nf := c.Get("fileUpload").(string)
	if !nf {
		file = ""
	}

	request := listdto.UpdateList{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
		Upload:    file,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.HandlerRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}
	if list.Upload != "" {
		os.Remove(list.Upload)
	}

	if request.Title != "" {
		list.Title = request.Title
	}
	if request.Deskripsi != "" {
		list.Deskripsi = request.Deskripsi
	}
	if request.Upload != "" {
		list.Upload = request.Upload
	}

	data, err := h.HandlerRepository.UpdateList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerList) DeleteList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.HandlerRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}
	if list.Upload != "" {
		os.Remove(list.Upload)
	}

	data, err := h.HandlerRepository.DeleteList(list.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
