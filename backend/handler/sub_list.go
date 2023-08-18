package handler

import (
	dto "Moonlay/dto/result"
	sublistdto "Moonlay/dto/sublist"
	"Moonlay/models"
	"Moonlay/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerSubList struct {
	HandlerRepository repository.SubList
}

func HandlerSubList(HandlerRepository repository.SubList) *handlerSubList {
	return &handlerSubList{HandlerRepository}
}

func (h *handlerSubList) GetSubList(c echo.Context) error {
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

	totalRecords, _ := h.HandlerRepository.GetTotalRecordsSubList()

	sublists, err := h.HandlerRepository.FindAll(page, totalRecords, pageSize, title, deskripsi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: sublists})
}

func (h *handlerSubList) GetSubByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	subList, err := h.HandlerRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: subList})
}

func (h *handlerSubList) CreateSubList(c echo.Context) error {
	file, nf := c.Get("fileUpload").(string)
	if !nf {
		file = ""
	}

	listId, _ := strconv.Atoi(c.FormValue("list_id"))

	request := sublistdto.CreateSubList{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
		Upload:    file,
		ListId:    listId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}

	list := models.SubList{
		Title:     request.Title,
		Deskripsi: request.Deskripsi,
		Upload:    request.Upload,
		ListId:    request.ListId,
	}

	data, err := h.HandlerRepository.CreateSubList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})

}

func (h *handlerSubList) UpdateSubList(c echo.Context) error {
	file, nf := c.Get("fileUpload").(string)
	if !nf {
		file = ""
	}

	listId, _ := strconv.Atoi(c.FormValue("list_id"))

	request := sublistdto.UpdateSubList{
		Title:     c.FormValue("title"),
		Deskripsi: c.FormValue("deskripsi"),
		Upload:    file,
		ListId:    listId,
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
	if request.ListId != 0 {
		list.ListId = request.ListId
	}

	data, err := h.HandlerRepository.UpdateSubList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})

}

func (h *handlerSubList) DeleteSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.HandlerRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
	}
	if list.Upload != "" {
		os.Remove(list.Upload)
	}

	data, err := h.HandlerRepository.DeleteSubList(list.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Data: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}

func convertResponse(subList models.SubList) sublistdto.ConvertResponse {
	return sublistdto.ConvertResponse{
		Id:        subList.Id,
		Title:     subList.Title,
		Deskripsi: subList.Deskripsi,
		Upload:    subList.Upload,
		ListId:    subList.ListId,
	}
}
