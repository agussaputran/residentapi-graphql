package responses

// OfficeResponse model
type OfficeResponse struct {
	ID          uint   `json:"id"`
	SubDistrict string `json:"sub_district"`
	Office      string `json:"office"`
}
