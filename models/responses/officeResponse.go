package responses

// OfficeResponse model
type OfficeResponse struct {
	ID              uint   `json:"id"`
	SubDistrictName string `json:"sub_district_name"`
	OfficeName      string `json:"office_name"`
}
