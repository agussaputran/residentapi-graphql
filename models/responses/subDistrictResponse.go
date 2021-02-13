package responses

// SubDistrictResponse model
type SubDistrictResponse struct {
	ID          uint   `json:"id"`
	Subdistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
}
