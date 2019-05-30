package base

import (
	"github.com/mochadwi/go-article/models"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// GetUUUID
func GetUUID(name string) string {
	return uuid.NewV5(uuid.Must(uuid.NewV4()), name).String()
}

// Check for request is valid
func IsRequestValid(m interface{}) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Get for status code based on received error
func GetStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
