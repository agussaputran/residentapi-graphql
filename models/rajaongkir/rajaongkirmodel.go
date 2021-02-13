package rajaongkir

// Prov model
type Prov struct {
	RajaOngkir ProvinceResponse `json:"rajaongkir"`
}

// ProvinceResponse model
type ProvinceResponse struct {
	ProvinceResults []ProvinceResult `json:"results"`
}

// ProvinceResult model
type ProvinceResult struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

// ============================================================================

// City model
type City struct {
	RajaOngkir CityResponse `json:"rajaongkir"`
}

// CityResponse model
type CityResponse struct {
	CityResults []CityResult `json:"results"`
}

// CityResult model
type CityResult struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}
