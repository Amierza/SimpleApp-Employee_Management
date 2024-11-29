package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/Amierza/employee-management/dto"
	"github.com/Amierza/employee-management/entity"
	"gorm.io/gorm"
)

type (
	EmployeeRepository interface {
		CreateEmployee(ctx context.Context, tx *gorm.DB, employee entity.Employee) (entity.Employee, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Employee, bool, error)
		CheckPhoneNumber(ctx context.Context, tx *gorm.DB, phoneNumber string) (entity.Employee, bool, error)
		FindJobByJobID(ctx context.Context, tx *gorm.DB, jobID string) (entity.Job, error)
		FindEmployeeByEmployeeID(ctx context.Context, tx *gorm.DB, employeeID string) (entity.Employee, error)
		GetAllEmployeeWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllEmployeeRepositoryResponse, error)
		UpdateEmployee(ctx context.Context, tx *gorm.DB, employee entity.Employee) error
		DeleteEmployee(ctx context.Context, tx *gorm.DB, employee entity.Employee) error
	}

	employeeRepository struct {
		db *gorm.DB
	}
)

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, tx *gorm.DB, employee entity.Employee) (entity.Employee, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&employee).Error; err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (r *employeeRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Employee, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var employee entity.Employee
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&employee).Error; err != nil {
		return entity.Employee{}, false, err
	}

	return employee, true, nil
}

func (r *employeeRepository) CheckPhoneNumber(ctx context.Context, tx *gorm.DB, phoneNumber string) (entity.Employee, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var employee entity.Employee
	if err := tx.WithContext(ctx).Where("phone_number = ?", phoneNumber).Take(&employee).Error; err != nil {
		return entity.Employee{}, false, err
	}

	return employee, true, nil
}

func (r *employeeRepository) FindJobByJobID(ctx context.Context, tx *gorm.DB, jobID string) (entity.Job, error) {
	if tx == nil {
		tx = r.db
	}

	var job entity.Job
	if err := tx.WithContext(ctx).Where("job_id = ?", jobID).Take(&job).Error; err != nil {
		return entity.Job{}, err
	}

	return job, nil
}

func (r *employeeRepository) FindEmployeeByEmployeeID(ctx context.Context, tx *gorm.DB, employeeID string) (entity.Employee, error) {
	if tx == nil {
		tx = r.db
	}

	var employee entity.Employee
	if err := tx.WithContext(ctx).Where("employee_id = ?", employeeID).Take(&employee).Error; err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (r *employeeRepository) GetAllEmployeeWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllEmployeeRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var employees []entity.Employee
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Employee{})

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllEmployeeRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&employees).Error; err != nil {
		return dto.GetAllEmployeeRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllEmployeeRepositoryResponse{
		Employees: employees,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *employeeRepository) UpdateEmployee(ctx context.Context, tx *gorm.DB, employee entity.Employee) error {
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Save(&employee).Error
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, tx *gorm.DB, employee entity.Employee) error {
	if tx == nil {
		tx = r.db
	}

	err := tx.WithContext(ctx).Unscoped().Delete(&employee).Error
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	return nil
}
