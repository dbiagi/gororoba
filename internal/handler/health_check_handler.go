package handler

type HealthCheckHandler struct {
}

const (
	HealthStatusUP   = "UP"
	HealthStatusDOWN = "DOWN"
)

type HealthCheckResponse struct {
	Status   string         `json:"status"`
	Web      WebStatus      `json:"web"`
	Database DatabaseStatus `json:"database"`
}

type WebStatus struct {
	Status string `json:"status"`
}

type DatabaseStatus struct {
	Status string `json:"status"`
}

func NewHealthCheckHandler() HealthCheckHandler {
	return HealthCheckHandler{}
}

func (h *HealthCheckHandler) Check() HealthCheckResponse {
	return HealthCheckResponse{Status: HealthStatusUP, Web: WebStatus{Status: HealthStatusUP}}
}

func (h *HealthCheckHandler) CheckComplete() HealthCheckResponse {
	return HealthCheckResponse{Status: HealthStatusUP, Web: WebStatus{Status: HealthStatusUP}, Database: DatabaseStatus{Status: HealthStatusUP}}
}
