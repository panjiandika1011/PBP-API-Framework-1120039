package Controller

type Patient struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
}

type PatientResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Patient `json:"data"`
}

type MedicalStaff struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

type MedicalStaffResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []MedicalStaff `json:"data"`
}

type Schedule struct {
	ID         int    `json:"id"`
	DateTime   string `json:"datetime"`
	PatientID  int    `json:"patientid"`
	MedstaffID int    `json:"medstaffid"`
}

type ScheduleResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Schedule `json:"data"`
}

type ScheduleDetail struct {
	ID       int          `json:"id"`
	DateTime string       `json:"datetime"`
	Patient  Patient      `json:"patient"`
	Medstaff MedicalStaff `json:"medstaff"`
}

type ScheduleDetailResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []ScheduleDetail `json:"data"`
}
