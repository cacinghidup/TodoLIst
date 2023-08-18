package middleware

import (
	dto "Moonlay/dto/result"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadFile(e echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			c.Set("fileUpload", "")
			return e(c)
		}

		ext := filepath.Ext(file.Filename)
		if ext != ".txt" && ext != ".pdf" {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: "File not supported"})
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
		}

		defer src.Close()

		tempFile, err := os.CreateTemp("./upload", "file-*"+ext+"")
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
		}

		defer tempFile.Close()

		_, err = io.Copy(tempFile, src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Data: err.Error()})
		}

		ctx := tempFile.Name()

		c.Set("fileUpload", ctx)

		return e(c)

	}
}
