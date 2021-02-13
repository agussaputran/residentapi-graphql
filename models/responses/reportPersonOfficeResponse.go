package responses

// ReportPersonOfficeResponse model
type ReportPersonOfficeResponse struct {
	ID       uint     `json:"id"`
	FullName string   `json:"full_name"`
	Total    int64    `json:"total"`
	CityName []string `json:"city_name"`
}
