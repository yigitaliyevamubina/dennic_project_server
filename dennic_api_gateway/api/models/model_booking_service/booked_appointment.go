package model_booking_service

type Appointment struct {
	Id              int64   `json:"id"`
	DepartmentId    string  `json:"department_id"`
	DoctorId        string  `json:"doctor_id"`
	PatientId       string  `json:"patient_id"`
	AppointmentDate string  `json:"appointment_date"`
	AppointmentTime string  `json:"appointment_time"`
	Duration        int64   `json:"duration"`
	Key             string  `json:"key"`
	ExpiresAt       string  `json:"expires_at"`
	PatientStatus   string  `json:"patient_status"`
	PatientProblem  string  `json:"patient_problem"`
	DoctorServiceId string  `json:"doctor_service_id"`
	PaymentType     string  `json:"payment_type"`
	PaymentAmount   float64 `json:"payment_amount"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type AppointmentsType struct {
	Count        int64          `json:"count"`
	Appointments []*Appointment `json:"appointments"`
}

type CreateAppointmentReq struct {
	DepartmentId    string  `json:"department_id"`
	DoctorId        string  `json:"doctor_id"`
	PatientId       string  `json:"patient_id"`
	AppointmentDate string  `json:"appointment_date"`
	AppointmentTime string  `json:"appointment_time"`
	Duration        int64   `json:"duration"`
	Key             string  `json:"key"`
	ExpiresAt       string  `json:"expires_at"`
	PatientStatus   string  `json:"patient_status"`
	PatientProblem  string  `json:"patient_problem"`
	DoctorServiceId string  `json:"doctor_service_id"`
	PaymentType     string  `json:"payment_type"`
	PaymentAmount   float64 `json:"payment_amount"`
}

type UpdateAppointmentReq struct {
	BookedAppointmentId string  `json:"booked_appointment_id"`
	AppointmentDate     string  `json:"appointment_date"`
	AppointmentTime     string  `json:"appointment_time"`
	Duration            int64   `json:"duration"`
	Key                 string  `json:"key"`
	ExpiresAt           string  `json:"expires_at"`
	PatientStatus       string  `json:"patient_status"`
	PatientProblem      string  `json:"patient_problem"`
	DoctorServiceId     string  `json:"doctor_service_id"`
	PaymentType         string  `json:"payment_type"`
	PaymentAmount       float64 `json:"payment_amount"`
}
