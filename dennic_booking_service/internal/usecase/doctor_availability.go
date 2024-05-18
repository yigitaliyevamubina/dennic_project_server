package usecase

import (
	"booking_service/internal/entity/doctor_availability"
	"booking_service/internal/pkg/otlp"
	"context"
	"time"
)

const (
	serviceNameDoctorAvailability = "DoctorAvailabilityService"
	spanNameDoctorAvailability    = "DoctorAvailabilityUsecase"
)

// BookedDoctorAvailabilityUseCase -.
type BookedDoctorAvailabilityUseCase struct {
	Repo       DoctorAvailability
	ctxTimeout time.Duration
}

// NewBookedDoctorAvailability -.
func NewBookedDoctorAvailability(r DoctorAvailability, ctxTimeout time.Duration) *BookedDoctorAvailabilityUseCase {
	return &BookedDoctorAvailabilityUseCase{
		Repo:       r,
		ctxTimeout: ctxTimeout,
	}
}

func (r *BookedDoctorAvailabilityUseCase) CreateDoctorAvailability(ctx context.Context, req *doctor_availability.CreateDoctorAvailability) (*doctor_availability.DoctorAvailability, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorAvailability, spanNameDoctorAvailability+"Create")
	span.End()

	return r.Repo.CreateDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) GetDoctorAvailability(ctx context.Context, req *doctor_availability.FieldValueReq) (*doctor_availability.DoctorAvailability, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorAvailability, spanNameDoctorAvailability+"Get")
	span.End()

	return r.Repo.GetDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) GetAllDoctorAvailability(ctx context.Context, req *doctor_availability.GetAllReq) (*doctor_availability.DoctorAvailabilityType, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorAvailability, spanNameDoctorAvailability+"List")
	span.End()

	return r.Repo.GetAllDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) UpdateDoctorAvailability(ctx context.Context, req *doctor_availability.UpdateDoctorAvailability) (*doctor_availability.DoctorAvailability, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorAvailability, spanNameDoctorAvailability+"Update")
	span.End()

	return r.Repo.UpdateDoctorAvailability(ctx, req)
}

func (r *BookedDoctorAvailabilityUseCase) DeleteDoctorAvailability(ctx context.Context, req *doctor_availability.FieldValueReq) (*doctor_availability.StatusRes, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameDoctorAvailability, spanNameDoctorAvailability+"Delete")
	span.End()

	return r.Repo.DeleteDoctorAvailability(ctx, req)
}
