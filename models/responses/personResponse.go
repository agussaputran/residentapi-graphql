package responses

// PersonResponse model
type PersonResponse struct {
	ID              uint
	Nip             string `json:"nip"`
	FullName        string `json:"full_name"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	BirthDate       string `json:"birth_date"`
	BirthPlace      string `json:"birth_place"`
	PhotoProfileURL string `json:"photo_profile_url"`
	Gender          string `json:"gender"`
	ZoneLocation    string `json:"zone_location"`
	SubDistrict     string `json:"sub_district"`
}
