package mock_handler

import (
	"errors"
	"net/http"

	customerrors "cortex/pkg/custom_errors"

	"github.com/google/uuid"
)

type DeleteCategory struct {
	Name           string
	ID             string
	MockReturnErr  error
	WantStatusCode int
}

func DeleteCategoryTestData() []DeleteCategory {
	return []DeleteCategory{
		{
			Name:           "success",
			ID:             uuid.NewString(),
			MockReturnErr:  nil,
			WantStatusCode: http.StatusOK,
		},
		{
			Name:           "category not found",
			ID:             uuid.NewString(),
			MockReturnErr:  customerrors.ErrCategoryNotFound,
			WantStatusCode: http.StatusNotFound,
		},
		{
			Name:           "internal server error",
			ID:             uuid.NewString(),
			MockReturnErr:  errors.New("internal server error"),
			WantStatusCode: http.StatusInternalServerError,
		},
	}
}
