package responses

// DistrictResponse model
type DistrictResponse struct {
	ID           uint   `json:"id"`
	DistrictName string `json:"district_name"`
	ProvinceName string `json:"province_name"`
}
