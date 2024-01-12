package boardhandler

import (
	"net/http"
	"strconv"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) createBoard(c echo.Context) error {
	var req param.UserBoardCreateRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	req.CreatorUserID = claims.UserID
	id, err := h.boardSvc.CreateBoard(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (h Handler) getAllBoardByWorkspaceID(c echo.Context) error {

	workspaceID, err := strconv.ParseUint(c.QueryParam("workspace_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": workspaceID,
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	resp, err := h.boardSvc.GetAllBoardByWorkspaceID(param.UserBoardGetAllRequest{
		UserID: claims.UserID,
		WorkspaceID: uint(workspaceID),
	})

	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) deleteBoardByID(c echo.Context) error {

	boardID, err := strconv.ParseUint(c.Param("board_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "you should pass board_id to remove",
		})
	}

	if err := h.boardSvc.DeleteBoardByID(param.UserBoardDeleteRequest{
		BoardID: uint(boardID),
		}); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
