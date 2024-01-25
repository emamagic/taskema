package orghandler

import (
	"net/http"
	"strconv"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) createOrganization(c echo.Context) error {
	var req param.UserOrganizationCreateRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	req.CreatorUserID = claims.UserID

	if err := h.orgValidation.CreateOrganization(req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	id, err := h.orgSvc.CreateOrganization(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (h Handler) getAllOrganization(c echo.Context) error {
	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)

	resp, err := h.orgSvc.GetAllOrganizationByID(param.UserOrganizationGetAllRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) deleteOrganizationByID(c echo.Context) error {
	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)

	orgID, err := strconv.ParseUint(c.Param("organization_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "you should pass organization_id to remove",
		})
	}

	req := param.UserOrganizationDeleteRequest{OrganizationID: uint(orgID), UserID: claims.UserID}
	if vErr := h.orgValidation.DeleteOrganization(req); vErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "organization id is not valid",
		})
	} 

	if err := h.orgSvc.DeleteOrganizationByID(req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
