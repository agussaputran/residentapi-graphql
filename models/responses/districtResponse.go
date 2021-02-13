package responses

// DistrictResponse model
type DistrictResponse struct {
	ID       uint   `json:"id"`
	District string `json:"district"`
	Province string `json:"province"`
}
