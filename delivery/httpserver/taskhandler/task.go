package taskhandler

import (
	"net/http"
	"strconv"
	"taskema/param"
	"taskema/service/authservice"

	"github.com/labstack/echo/v4"
)

func (h Handler) createTask(c echo.Context) error {
	var req param.UserTaskCreateRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	req.CreatorUserID = claims.UserID
	id, err := h.taskservice.CreateTask(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (h Handler) getAllTaskByColumnID(c echo.Context) error {

	columnID, err := strconv.ParseUint(c.QueryParam("column_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	claims := c.Get(h.authCfg.ContextKey).(*authservice.Claims)
	resp, err := h.taskservice.GetAllTaskByColumnID(param.UserTaskGetAllRequest{
		ColumnID: uint(columnID),
		UserID:   claims.UserID,
	})
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) deleteTaskByID(c echo.Context) error {

	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 0)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "you should pass task_id to remove",
		})
	}

	if err := h.taskservice.DeleteTaskByID(param.UserTaskDeleteRequest{TaskID: uint(taskID)}); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
