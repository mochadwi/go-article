package httpecho

import (
	"context"
	"encoding/json"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/mochadwi/go-article/features/rating/mocks"
	"github.com/mochadwi/go-article/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHttpRatingHandler_GetByID_InvalidLessonID(t *testing.T) {
	e := echo.New()

	req, _ := http.NewRequest(echo.GET, "/rating/a", strings.NewReader(""))

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("rating/:lessonId")
	c.SetParamNames("lessonId")
	c.SetParamValues("a")

	handler := HttpRatingHandler{
		AUsecase: new(mocks.RatingUsecase),
	}

	handler.GetByID(c)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestHttpRatingHandler_GetByID_Succeed(t *testing.T) {
	var mockRating models.Rating
	err := faker.FakeData(&mockRating)
	assert.NoError(t, err)

	mockUCase := new(mocks.RatingUsecase)

	num := int(mockRating.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockRating, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/rating/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctx := c.Request().Context()
	assert.Equal(t, ctx, context.Background())

	c.SetPath("rating/:lessonId")
	c.SetParamNames("lessonId")
	c.SetParamValues(strconv.Itoa(num))
	handler := HttpRatingHandler{
		AUsecase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestHttpRatingHandler_Create_BadRequest(t *testing.T) {
	tempMockRating := models.Rating{
		ID:           0,
		LessonID:     1,
		RatingNumber: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	mockUCase := new(mocks.RatingUsecase)

	// todo error?
	//mockUCase.On("Create",
	//	mock.Anything,
	//	mock.AnythingOfType("*models.Rating")).Return(nil).Once()

	j, err := json.Marshal(tempMockRating)
	assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/rating", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/rating")

	handler := HttpRatingHandler{AUsecase: mockUCase}
	err = handler.Create(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	mockUCase.AssertExpectations(t)
}
