package responses

// SubDistrictResponse model
type SubDistrictResponse struct {
	ID              uint   `json:"id"`
	SubdistrictName string `json:"sub_district_name"`
	DistrictName    string `json:"district_name"`
	ProvinceName    string `json:"province_name"`
}
