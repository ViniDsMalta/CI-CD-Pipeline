package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ViniDsMalta/CI-CD-Pipeline/internal/handlers"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	rr := httptest.NewRecorder()

	handlers.HealthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf(
			"expect status code %d, received %d",
			http.StatusOK,
			rr.Code,
		)
	}
}