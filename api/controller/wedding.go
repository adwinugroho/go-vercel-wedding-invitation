package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/service"
	"github.com/labstack/echo/v4"
)

type WeddingController struct {
	WishesService service.WishesInterface
}

func NewController(wishesSvc service.WishesInterface) WeddingController {
	return WeddingController{
		WishesService: wishesSvc,
	}
}

func (h *WeddingController) GetListWishes(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error! Invalid number format on limit",
		})
	}

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error! Invalid number format on page",
		})
	}

	resp, err := h.WishesService.List(pageInt, limitInt)
	if err != nil {
		log.Printf("[controller] Error while get list:%+v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Internal Server Error, Please Contact Customer Service.",
		})
	} else if resp == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"message": "Data Not Found.",
			"data":    make([]string, 0),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"data":    resp,
		"message": "Success!",
	})
}
