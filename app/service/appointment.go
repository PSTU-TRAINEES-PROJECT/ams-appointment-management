package service

import (
	"ams-appointment-management/app/domain/models"
	"ams-appointment-management/app/domain/repository"
	"ams-appointment-management/app/serializer"
	"fmt"

	"ams-appointment-management/app/common/logger"
	"ams-appointment-management/app/common/pagination"
)

type AppointmentServiceInterface interface {
	Create(req *serializer.Appointment) (*models.Appointment, error)
	FindAll(pagination pagination.Page) (*[]models.Appointment, int64, error)
	FindByID(ID int) (*models.Appointment, error)
	Update(ID int, req *serializer.Appointment) (*models.Appointment, error)
	Delete(ID int) (int, error)
}

type AppointmentService struct {
	appointmentRepository repository.AppointmentRepositoryInterface
}

func NewAppointmentService(appointmentRepo repository.AppointmentRepositoryInterface) AppointmentService {
	return AppointmentService{appointmentRepository: appointmentRepo}
}

func (svc AppointmentService) Create(req *serializer.Appointment) (*models.Appointment, error) {
	appointment := models.Appointment{}
	appointment.ToUserModel(req)
	variant, err := svc.appointmentRepository.Create(&appointment)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Error:%s", err)
	}
	return variant, nil
}

func (svc AppointmentService) FindAll(pagination pagination.Page) (*[]models.Appointment, int64, error) {
	variants, count, err := svc.appointmentRepository.FindAll(pagination)
	if err != nil {
		logger.Error(err)
		return nil, 0, fmt.Errorf("Error:%s", err)
	}
	return variants, count, nil
}

func (svc AppointmentService) FindByID(ID int) (*models.Appointment, error) {
	appointment, err := svc.appointmentRepository.FindBy("id", ID)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Error:%s", err)
	}
	return appointment, nil
}

func (svc AppointmentService) Update(ID int, req *serializer.Appointment) (*models.Appointment, error) {
	user, err := svc.appointmentRepository.FindBy("id", ID)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Error:%s", err)
	}
	user.ToUserModel(req)
	user, err = svc.appointmentRepository.Update(user)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Error:%s", err)
	}
	return user, nil
}

func (svc AppointmentService) Delete(ID int) (int, error) {
	_, err := svc.appointmentRepository.FindBy("id", ID)
	if err != nil {
		logger.Error(err)
		return 0, fmt.Errorf("Error:%s", err)
	}
	ID, err = svc.appointmentRepository.Delete(ID)
	if err != nil {
		logger.Error(err)
		return 0, fmt.Errorf("Error:%s", err)
	}
	return ID, nil
}
