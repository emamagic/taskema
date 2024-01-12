package workspacehandler

import (
	"net/http"
	"strconv"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) createWorkspace(c echo.Context) error {
	var req param.UserWorkspaceCreateRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	req.CreatorUserID = claims.UserID

	if err := h.workspaceValidation.CreateWorkspace(req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	id, err := h.workspaceservice.CreateWorkspace(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (h Handler) getAllWorkspaceByOrganizationID(c echo.Context) error {

	organizationID := c.QueryParam("organization_id")
	if organizationID == "" {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "you shoud specify organization_id",
		})
	}

	orgID, pErr := strconv.ParseUint(organizationID, 10, 0)
	if pErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": pErr.Error(),
		})
	}

	req := param.UserWorkspaceGetAllRequest{OrganizationID: uint(orgID)}

	if err := h.workspaceValidation.GetAllWorkspace(req); err != nil {
		
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	resp, err := h.workspaceservice.GetAllWorkspaceByOrganizationID(req, claims.UserID)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) deleteWorkspaceByID(c echo.Context) error {

	workspaceID, err := strconv.ParseUint(c.Param("workspace_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "you should pass workspace_id to remove",
		})
	}

	if err := h.workspaceservice.DeleteWorkspaceByID(param.UserWorkspaceDeleteRequest{WorkspaceID: uint(workspaceID)}); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
