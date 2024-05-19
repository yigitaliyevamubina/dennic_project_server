package v1

import (
	"context"
	e "dennic_admin_api_gateway/api/handlers/regtool"
	"dennic_admin_api_gateway/api/models"
	"dennic_admin_api_gateway/api/models/model_healthcare_service"
	pb "dennic_admin_api_gateway/genproto/healthcare-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateDoctorService ...
// @Summary CreateDoctorService
// @Description CreateDoctorService - Api for crete doctor_services
// @Tags Doctor Services
// @Accept json
// @Produce json
// @Param DoctorServiceReq body model_healthcare_service.DoctorServicesReq true "DoctorServiceReq"
// @Success 200 {object} model_healthcare_service.DoctorServicesRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-services [post]
func (h *HandlerV1) CreateDoctorService(c *gin.Context) {
	var (
		body        model_healthcare_service.DoctorServicesReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateDoctorService") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorServices, err := h.serviceManager.HealthcareService().DoctorsService().CreateDoctorServices(ctx, &pb.DoctorServices{
		Id:               uuid.NewString(),
		DoctorId:         body.DoctorId,
		SpecializationId: body.SpecializationId,
		OnlinePrice:      body.OnlinePrice,
		OfflinePrice:     body.OfflinePrice,
		Name:             body.Name,
		Duration:         body.Duration,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateDoctorService") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorServicesRes{
		Id:               doctorServices.Id,
		Order:            doctorServices.DoctorServiceOrder,
		DoctorId:         doctorServices.DoctorId,
		SpecializationId: doctorServices.SpecializationId,
		OnlinePrice:      doctorServices.OnlinePrice,
		OfflinePrice:     doctorServices.OfflinePrice,
		Name:             doctorServices.Name,
		Duration:         doctorServices.Duration,
		CreatedAt:        doctorServices.CreatedAt,
		UpdatedAt:        e.UpdateTimeFilter(doctorServices.UpdatedAt),
	})
}

// GetDoctorService ...
// @Summary GetDoctorService
// @Description GetDoctorService - Api for get doctor_services
// @Tags Doctor Services
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} model_healthcare_service.DoctorServicesRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-services/get [get]
func (h *HandlerV1) GetDoctorService(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorServices, err := h.serviceManager.HealthcareService().DoctorsService().GetDoctorServiceByID(ctx, &pb.GetReqStr{
		Field:    "id",
		Value:    id,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetDoctorService") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorServicesRes{
		Id:               doctorServices.Id,
		Order:            doctorServices.DoctorServiceOrder,
		DoctorId:         doctorServices.DoctorId,
		SpecializationId: doctorServices.SpecializationId,
		OnlinePrice:      doctorServices.OnlinePrice,
		OfflinePrice:     doctorServices.OfflinePrice,
		Name:             doctorServices.Name,
		Duration:         doctorServices.Duration,
		CreatedAt:        doctorServices.CreatedAt,
		UpdatedAt:        e.UpdateTimeFilter(doctorServices.UpdatedAt),
	})
}

// ListDoctorServices ...
// @Summary ListDoctorServices
// @Description ListDoctorServices - Api for list doctor_services
// @Tags Doctor Services
// @Accept json
// @Produce json
// @Param ListReq query model_healthcare_service.ListReqDoctorServices false "ListReq"
// @Success 200 {object} model_healthcare_service.ListDoctorServices
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-services [get]
func (h *HandlerV1) ListDoctorServices(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListDoctorServices") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorServicess, err := h.serviceManager.HealthcareService().DoctorsService().GetAllDoctorServices(ctx, &pb.GetAllDoctorServiceS{
		Field:    "",
		Value:    "",
		IsActive: false,
		Page:     int64(pageInt),
		Limit:    int64(limitInt),
		OrderBy:  orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "ListDoctorServices") {
		return
	}
	var doctorServicessRes model_healthcare_service.ListDoctorServices
	for _, doctorServicesRes := range doctorServicess.DoctorServices {
		doctorServicessRes.DoctorServices = append(doctorServicessRes.DoctorServices, &model_healthcare_service.DoctorServicesRes{
			Id:               doctorServicesRes.Id,
			Order:            doctorServicesRes.DoctorServiceOrder,
			DoctorId:         doctorServicesRes.DoctorId,
			SpecializationId: doctorServicesRes.SpecializationId,
			OnlinePrice:      doctorServicesRes.OnlinePrice,
			OfflinePrice:     doctorServicesRes.OfflinePrice,
			Name:             doctorServicesRes.Name,
			Duration:         doctorServicesRes.Duration,
			CreatedAt:        doctorServicesRes.CreatedAt,
			UpdatedAt:        e.UpdateTimeFilter(doctorServicesRes.UpdatedAt),
		})
	}

	c.JSON(http.StatusOK, model_healthcare_service.ListDoctorServices{
		Count:          doctorServicess.Count,
		DoctorServices: doctorServicessRes.DoctorServices,
	})
}

// UpdateDoctorServices ...
// @Summary UpdateDoctorServices
// @Description UpdateDoctorServices - Api for update doctor_services
// @Tags Doctor Services
// @Accept json
// @Produce json
// @Param UpdateDoctorServicesReq body model_healthcare_service.DoctorServicesReq true "UpdateDoctorServicesReq"
// @Success 200 {object} model_healthcare_service.DoctorServicesRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-services [put]
func (h *HandlerV1) UpdateDoctorServices(c *gin.Context) {
	var (
		body        model_healthcare_service.DoctorServicesReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateDoctorServices") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	doctorServices, err := h.serviceManager.HealthcareService().DoctorsService().UpdateDoctorServices(ctx, &pb.DoctorServices{
		Id:               body.Id,
		DoctorId:         body.DoctorId,
		SpecializationId: body.SpecializationId,
		OnlinePrice:      body.OnlinePrice,
		OfflinePrice:     body.OfflinePrice,
		Name:             body.Name,
		Duration:         body.Duration,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateDoctorServices") {
		return
	}

	c.JSON(http.StatusOK, model_healthcare_service.DoctorServicesRes{
		Id:               doctorServices.Id,
		Order:            doctorServices.DoctorServiceOrder,
		DoctorId:         doctorServices.DoctorId,
		SpecializationId: doctorServices.SpecializationId,
		OnlinePrice:      doctorServices.OnlinePrice,
		OfflinePrice:     doctorServices.OfflinePrice,
		Name:             doctorServices.Name,
		Duration:         doctorServices.Duration,
		CreatedAt:        doctorServices.CreatedAt,
		UpdatedAt:        e.UpdateTimeFilter(doctorServices.UpdatedAt),
	})
}

// DeleteDoctorService ...
// @Summary DeleteDoctorService
// @Description DeleteDoctorService - Api for delete doctorServices
// @Tags Doctor Services
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/doctor-services [delete]
func (h *HandlerV1) DeleteDoctorService(c *gin.Context) {
	value := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	status, err := h.serviceManager.HealthcareService().DoctorsService().DeleteDoctorService(ctx, &pb.GetReqStr{
		Field:    "id",
		Value:    value,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteDoctorService") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: status.Status})
}
