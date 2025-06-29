package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/pkg/helpers"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/service"
	"github.com/labstack/echo/v4"
)

type WeddingController struct {
	WishesService service.WishesInterface
	RSVPService   service.RSVPInterface
	GiftService   service.GiftInterface
}

func NewController(wishesSvc service.WishesInterface, rsvpSvc service.RSVPInterface, giftSvc service.GiftInterface) WeddingController {
	return WeddingController{
		WishesService: wishesSvc,
		RSVPService:   rsvpSvc,
		GiftService:   giftSvc,
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
		log.Printf("[controller] Error while get list wishes:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
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

func (h *WeddingController) NewWishes(c echo.Context) error {
	var body model.RequestNewWish
	err := c.Bind(&body)
	if err != nil {
		log.Println("[NewWishes] Error cause:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error binding request body",
		})
	}

	if body.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error name required.",
		})
	}

	if body.Message == "" || len(body.Message) < 4 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error message required",
		})
	}

	resp, err := h.WishesService.New(model.Wishes{
		Name:        body.Name,
		Message:     body.Message,
		IsPublished: true,
		CreatedAt:   helpers.TimeHostNow("Asia/Jakarta").Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Printf("[controller] Error create a new wish:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Internal Server Error, Please Contact Customer Service.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"data":    resp,
		"message": "Success!",
	})
}

func (h *WeddingController) GetListAttending(c echo.Context) error {
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

	isAttending := c.QueryParam("is_attending")
	isAttendingBoolean := false
	if isAttending == "" {
		isAttendingBoolean = false
	}

	if isAttending == "false" {
		isAttendingBoolean = false
	} else if isAttending == "true" {
		isAttendingBoolean = true
	}

	resp, err := h.RSVPService.List(pageInt, limitInt, isAttendingBoolean)
	if err != nil {
		log.Printf("[controller] Error while get list rsvp:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
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

func (h *WeddingController) NewReservation(c echo.Context) error {
	var body model.RequestNewRSVP
	err := c.Bind(&body)
	if err != nil {
		log.Println("[NewReservation] Error cause:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error binding request body",
		})
	}

	if body.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error name required.",
		})
	}

	if body.GuestCount > 8 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Guest count is reached limit",
		})
	}

	if body.GuestCount == 0 {
		body.GuestCount = 1
	}

	initAttending := false
	if body.IsAttending == nil {
		body.IsAttending = &initAttending
	}

	resp, err := h.RSVPService.New(model.Reservation{
		Name:        body.Name,
		GuestCount:  body.GuestCount,
		IsAttending: *body.IsAttending,
		CreatedAt:   helpers.TimeHostNow("Asia/Jakarta").Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Printf("[controller] Error while create new reservation:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Internal Server Error, Please Contact Customer Service.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"data":    resp,
		"message": "Success!",
	})
}

func (h *WeddingController) GetListGift(c echo.Context) error {
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

	resp, err := h.GiftService.List(pageInt, limitInt)
	if err != nil {
		log.Printf("[controller] Error while get list gift:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
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

func (h *WeddingController) NewGift(c echo.Context) error {
	var body model.RequestNewGift
	err := c.Bind(&body)
	if err != nil {
		log.Println("[NewWishes] Error cause:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error binding request body",
		})
	}

	if body.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error name required.",
		})
	}

	if body.Price == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error price required.",
		})
	}

	if body.StoreLink == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error store link required.",
		})
	}

	resp, err := h.GiftService.New(model.Gifts{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		ImageURL:    body.ImageURL,
		Category:    body.Category,
		StoreLink:   body.StoreLink,
		CreatedAt:   helpers.TimeHostNow("Asia/Jakarta").Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Printf("[controller] Error create a new gift:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Internal Server Error, Please Contact Customer Service.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"data":    resp,
		"message": "Success!",
	})
}

func (h *WeddingController) MarkedGiftAsBooked(c echo.Context) error {
	var body model.RequestMarkedGiftAsBooked
	err := c.Bind(&body)
	if err != nil {
		log.Println("[NewWishes] Error cause:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error binding request body",
		})
	}

	if body.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error name required.",
		})
	}

	if body.IsBooked == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Error price required.",
		})
	}

	nowStr := helpers.TimeHostNow("Asia/Jakarta").Format("2006-01-02 15:04:05")
	err = h.GiftService.MarkedGiftAsBooked(model.Gifts{
		ID:        body.ID,
		IsBooked:  *body.IsBooked,
		UpdatedAt: &nowStr,
	})
	if err != nil {
		log.Printf("[controller] Error updating gift to booked:%+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Internal Server Error, Please Contact Customer Service.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"data":    body.ID,
		"message": "Success!",
	})
}
