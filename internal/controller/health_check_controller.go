package controller

import (
	"gororoba/internal/handler"
	"net/http"
)

type HealthCheckController struct {
	handler.HealthCheckHandler
}

func NewHealthCheckController(h handler.HealthCheckHandler) HealthCheckController {
	return HealthCheckController{HealthCheckHandler: h}
}

func (h *HealthCheckController) Check(w http.ResponseWriter, r *http.Request) HttpResponse {
	health := h.HealthCheckHandler.Check()

	if health.Status == handler.HealthStatusDOWN {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return HttpResponse{Body: health}
}

func (h *HealthCheckController) CheckComplete(w http.ResponseWriter, r *http.Request) HttpResponse {
	health := h.HealthCheckHandler.CheckComplete()

	var statusCode = http.StatusOK
	if health.Status == handler.HealthStatusDOWN {
		statusCode = http.StatusInternalServerError
	}

	return HttpResponse{StatusCode: statusCode, Body: health}
}
