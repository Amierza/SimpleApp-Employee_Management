package service

import (
	"context"
	"sync"

	"github.com/Amierza/employee-management/dto"
	"github.com/Amierza/employee-management/entity"
	"github.com/Amierza/employee-management/repository"
	"github.com/google/uuid"
)

type (
	EmployeeService interface {
		CreateEmployee(ctx context.Context, req dto.EmployeeCreateRequest) (dto.EmployeeCreateResponse, error)
		GetAllEmployeeWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.EmployeePaginationResponse, error)
		UpdateProfileEmployee(ctx context.Context, req dto.UpdateProfileRequest) (dto.EmployeeCreateResponse, error)
		DeleteEmployee(ctx context.Context, req dto.DeleteEmployeeRequest) error
	}

	employeeService struct {
		employeeRepo repository.EmployeeRepository
	}
)

var (
	mu sync.Mutex
)

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepo: employeeRepo,
	}
}

func (s *employeeService) CreateEmployee(ctx context.Context, req dto.EmployeeCreateRequest) (dto.EmployeeCreateResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag, err := s.employeeRepo.CheckEmail(ctx, nil, req.Email)
	if err == nil || flag {
		return dto.EmployeeCreateResponse{}, dto.ErrEmployeeIDAlreadyExists
	}

	job, err := s.employeeRepo.FindJobByJobID(ctx, nil, req.JobID)
	if err != nil {
		return dto.EmployeeCreateResponse{}, dto.ErrJobNotFound
	}

	employee := entity.Employee{
		EmployeeID:  uuid.New(),
		JobID:       job.JobID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Salary:      req.Salary,
	}

	newEmployee, err := s.employeeRepo.CreateEmployee(ctx, nil, employee)
	if err != nil {
		return dto.EmployeeCreateResponse{}, dto.ErrCreateEmployee
	}

	return dto.EmployeeCreateResponse{
		EmployeeID:  newEmployee.EmployeeID.String(),
		JobID:       newEmployee.JobID.String(),
		FirstName:   newEmployee.FirstName,
		LastName:    newEmployee.LastName,
		Email:       newEmployee.Email,
		PhoneNumber: newEmployee.PhoneNumber,
		Salary:      newEmployee.Salary,
	}, nil
}

func (s *employeeService) GetAllEmployeeWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.EmployeePaginationResponse, error) {
	dataWithPaginate, err := s.employeeRepo.GetAllEmployeeWithPagination(ctx, nil, req)
	if err != nil {
		return dto.EmployeePaginationResponse{}, err
	}

	var datas []dto.AllEmployeeResponse
	for _, employee := range dataWithPaginate.Employees {
		data := dto.AllEmployeeResponse{
			EmployeeID:  employee.EmployeeID.String(),
			JobID:       employee.JobID.String(),
			FirstName:   employee.FirstName,
			LastName:    employee.LastName,
			Email:       employee.Email,
			PhoneNumber: employee.PhoneNumber,
			Salary:      employee.Salary,
		}

		datas = append(datas, data)
	}

	return dto.EmployeePaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *employeeService) UpdateProfileEmployee(ctx context.Context, req dto.UpdateProfileRequest) (dto.EmployeeCreateResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	employee, err := s.employeeRepo.FindEmployeeByEmployeeID(ctx, nil, req.EmployeeID)
	if err != nil {
		return dto.EmployeeCreateResponse{}, dto.ErrGetEmployeeFromEmployeeID
	}

	if req.FirstName == "" {
		req.FirstName = employee.FirstName
	}

	if req.LastName == "" {
		req.LastName = employee.LastName
	}

	if req.PhoneNumber == "" {
		req.PhoneNumber = employee.PhoneNumber
	} else {
		_, flag, err := s.employeeRepo.CheckPhoneNumber(ctx, nil, req.PhoneNumber)
		if err == nil || flag {
			return dto.EmployeeCreateResponse{}, dto.ErrPhoneNumberAlreadyExists
		}
	}

	job, err := s.employeeRepo.FindJobByJobID(ctx, nil, req.JobID)
	if err != nil {
		return dto.EmployeeCreateResponse{}, dto.ErrJobNotFound
	}

	if req.Email == "" {
		req.Email = employee.Email
	} else {
		_, flag, err := s.employeeRepo.CheckEmail(ctx, nil, req.Email)
		if err == nil || flag {
			return dto.EmployeeCreateResponse{}, dto.ErrEmailAlreadyExists
		}
	}

	if req.Salary == 0 {
		req.Salary = employee.Salary
	}

	updatedEmployee := entity.Employee{
		EmployeeID:  employee.EmployeeID,
		JobID:       job.JobID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Salary:      req.Salary,
	}

	if err := s.employeeRepo.UpdateEmployee(ctx, nil, updatedEmployee); err != nil {
		return dto.EmployeeCreateResponse{}, dto.ErrUpdateEmployee
	}

	return dto.EmployeeCreateResponse{
		EmployeeID:  updatedEmployee.EmployeeID.String(),
		JobID:       updatedEmployee.JobID.String(),
		FirstName:   updatedEmployee.FirstName,
		LastName:    updatedEmployee.LastName,
		Email:       updatedEmployee.Email,
		PhoneNumber: updatedEmployee.PhoneNumber,
		Salary:      updatedEmployee.Salary,
	}, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, req dto.DeleteEmployeeRequest) error {
	mu.Lock()
	defer mu.Unlock()

	employee, err := s.employeeRepo.FindEmployeeByEmployeeID(ctx, nil, req.EmployeeID)
	if err != nil {
		return dto.ErrGetEmployeeFromEmployeeID
	}

	if err := s.employeeRepo.DeleteEmployee(ctx, nil, employee); err != nil {
		return dto.ErrDeleteEmployee
	}

	return nil
}
