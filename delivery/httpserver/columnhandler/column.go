package columnhandler

import (
	"fmt"
	"net/http"
	"strconv"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) createColumn(c echo.Context) error {
	var req param.UserColumnCreateRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	req.CreatorUserID = claims.UserID
	id, err := h.columnSvc.CreateColumn(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (h Handler) getAllColumnByWorkspaceID(c echo.Context) error {

	workspaceID, err := strconv.ParseUint(c.QueryParam("workspace_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": fmt.Sprintf("there is no such workspace_id like: %d", workspaceID),
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	resp, err := h.columnSvc.GetAllColumnByWorkspaceID(param.UserColumnGetAllRequest{
		UserID:      claims.UserID,
		WorkspaceID: uint(workspaceID),
	})

	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) deleteColumnByID(c echo.Context) error {
	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)

	columnID, err := strconv.ParseUint(c.Param("column_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "you should pass column_id to remove",
		})
	}

	req := param.UserColumnDeleteRequest{
		ColumnID: uint(columnID),
		UserID:   claims.UserID,
	}
	if err := h.columnSvc.DeleteColumnByID(req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
