package middleware_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizeMultipleRoles(t *testing.T) {
	var handler = func(w http.ResponseWriter, r *http.Request) {

	}
	guardedHandler := middleware.AuthorizeMultipleRoles(handler, []int{businesslogic.AccountTypeAthlete, businesslogic.AccountTypeAdjudicator, businesslogic.AccountTypeScrutineer})

	req, _ := http.NewRequest(http.MethodPost, "/api/v1.0/account/authenticate", nil)
	req.Header.Set("Authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsaWNlLmFuZGVyc29uQGVtYWlsLmNvbSIsImV4cCI6MTUzMjM5NzEyNCwiaXNzdWVkIjoxNTMyMTM3OTI0LCJuYW1lIjoiQWxpY2UgQW5kZXJzb24iLCJ1dWlkIjoiYjZiM2UzMTUtZWNmMC00YmU1LWEzNzMtNmVlZmFjMmZmODQ2In0.1zrKzTMMXiCXBLVC63N0gbOtM5C9eknbeFjSEkjRgyg")

	recorder := httptest.NewRecorder()
	guardedHandler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}
