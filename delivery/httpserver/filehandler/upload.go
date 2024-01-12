package filehandler

import (
	"net/http"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

// TODO - File Upload Security Measures
func (h Handler) uploadFile(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	files := form.File["file"]
	req := param.FileUploadRequest{
		Files:         files,
		UserCreatorID: claims.UserID,
	}
	resp, sErr := h.fileSvc.StoreFile(req)
	if sErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": sErr.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)

}
