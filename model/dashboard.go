package model

type DashboardResponse struct {
	Products []*ProductResponse          `json:"product"`
	Services []*ServiceResponseDashboard `json:"service"`
}
