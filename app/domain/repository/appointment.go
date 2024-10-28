package repository

import (
	"ams-appointment-management/app/common/pagination"
	"ams-appointment-management/app/domain/models"

	"fmt"

	"gorm.io/gorm"
)

type AppointmentRepositoryInterface interface {
	Create(user *models.Appointment) (*models.Appointment, error)
	FindAll(pagination pagination.Page) (*[]models.Appointment, int64, error)
	FindBy(field string, value any) (*models.Appointment, error)
	Update(appointment *models.Appointment) (*models.Appointment, error)
	Delete(ID int) (int, error)
}

type AppointmentRepository struct {
	Db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return AppointmentRepository{Db: db}
}

func (repo AppointmentRepository) Create(appointment *models.Appointment) (*models.Appointment, error) {
	if err := repo.Db.Table("appointments").Create(&appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (repo AppointmentRepository) FindAll(pagination pagination.Page) (*[]models.Appointment, int64, error) {
	appointment := []models.Appointment{}
	count := new(int64)
	if err := repo.Db.Table("appointments").Count(count).Error; err != nil {
		return nil, 0, err
	}
	if err := repo.Db.Table("appointments").Offset(*pagination.Offset).Limit(*pagination.Limit).Order("id asc,id").Find(&appointment).Error; err != nil {
		return nil, 0, err
	}
	return &appointment, *count, nil
}

func (repo AppointmentRepository) FindBy(field string, value any) (user *models.Appointment, err error) {
	query := fmt.Sprintf("%s = ?", field)
	if err := repo.Db.Table("appointments").Where(query, value).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo AppointmentRepository) Update(appointment *models.Appointment) (*models.Appointment, error) {
	if err := repo.Db.Table("appointments").Save(&appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (repo AppointmentRepository) Delete(ID int) (int, error) {
	appointment := models.Appointment{}
	if err := repo.Db.Table("appointments").Where("id = ?", ID).Delete(&appointment).Unscoped().Error; err != nil {
		return 0, err
	}
	return ID, nil
}
