package httpecho

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mochadwi/go-article/models"

	"github.com/labstack/echo"
	baseHandler "github.com/mochadwi/go-article/base"
	ratingUcase "github.com/mochadwi/go-article/features/rating"
	"time"
)

type HttpRatingHandler struct {
	AUsecase ratingUcase.RatingUsecase
}

func (a *HttpRatingHandler) GetByID(c echo.Context) error {

	idParam, err := strconv.Atoi(c.Param("lessonId"))
	lessonId := int64(idParam)

	var response = &models.BaseResponse{
		RequestID: baseHandler.GetUUID(string(lessonId)),
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())

		return c.JSON(response.Code, response)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByID(ctx, lessonId)

	if err != nil {
		response.Code = baseHandler.GetStatusCode(err)
		response.Message = string(err.Error())

		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusOK
	response.Message = models.DATA_CREATED_SUCCESS
	response.Data = art
	return c.JSON(response.Code, art)
}

func (a *HttpRatingHandler) Create(c echo.Context) error {
	var rating models.Rating
	err := c.Bind(&rating)

	var response = &models.BaseResponse{
		RequestID: baseHandler.GetUUID(string(rating.ID)),
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		return c.JSON(response.Code, response)
	}

	if ok, err := baseHandler.IsRequestValid(&rating); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		return c.JSON(response.Code, response)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Create(ctx, &rating)

	if err != nil {
		response.Code = baseHandler.GetStatusCode(err)
		response.Message = string(err.Error())
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusCreated
	response.Message = models.DATA_CREATED_SUCCESS
	response.Data = ar
	return c.JSON(response.Code, response)
}

func NewRatingHttpEchoHandler(e *echo.Echo, us ratingUcase.RatingUsecase) {
	handler := &HttpRatingHandler{
		AUsecase: us,
	}

	e.GET("/rating/:lessonId", handler.GetByID)
	e.POST("/rating", handler.Create)
}
