package responses

// ReportCountPersonOfficeResponse model
type ReportCountPersonOfficeResponse struct {
	OfficeName string `json:"office_name"`
	CounPerson int64  `json:"count_person"`
}
