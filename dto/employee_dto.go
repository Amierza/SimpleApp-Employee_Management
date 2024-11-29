package dto

import (
	"errors"

	"github.com/Amierza/employee-management/entity"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_CREATE_EMPLOYEE         = "failed to create employee"
	MESSAGE_FAILED_GET_LIST_EMPLOYEE       = "failed to get list employee"
	MESSAGE_FAILED_UPDATE_PROFILE_EMPLOYEE = "failet update profile employee"
	MESSAGE_FAILED_DELETE_EMPLOYEE         = "failet delete employee"

	// Success
	MESSAGE_SUCCESS_CREATE_EMPLOYEE         = "success to create employee"
	MESSAGE_SUCCESS_GET_LIST_EMPLOYEE       = "success to get list employee"
	MESSAGE_SUCCESS_UPDATE_PROFILE_EMPLOYEE = "success update profile employee"
	MESSAGE_SUCCESS_DELETE_EMPLOYEE         = "success delete employee"
)

var (
	ErrCreateEmployee            = errors.New("failed to create employee")
	ErrEmployeeIDAlreadyExists   = errors.New("employee id already exists")
	ErrJobNotFound               = errors.New("job not found")
	ErrGetEmployeeFromEmployeeID = errors.New("failed to get employee by employee id")
	ErrPhoneNumberAlreadyExists  = errors.New("phone number already exists")
	ErrEmailAlreadyExists        = errors.New("email already exists")
	ErrUpdateEmployee            = errors.New("failed to update employee")
	ErrDeleteEmployee            = errors.New("failed to delete employee")
)

type (
	EmployeeCreateRequest struct {
		JobID       string  `json:"job_id" form:"job_id"`
		FirstName   string  `json:"first_name" form:"first_name"`
		LastName    string  `json:"last_name" form:"last_name"`
		Email       string  `json:"email" form:"email"`
		PhoneNumber string  `json:"phone_number" form:"phone_number"`
		Salary      float64 `json:"salary" form:"salary"`
	}

	EmployeeCreateResponse struct {
		EmployeeID  string  `json:"employee_id"`
		JobID       string  `json:"job_id"`
		FirstName   string  `json:"first_name"`
		LastName    string  `json:"last_name"`
		Email       string  `json:"email"`
		PhoneNumber string  `json:"phone_number"`
		Salary      float64 `json:"salary"`
		entity.TimeStamp
	}

	AllEmployeeResponse struct {
		EmployeeID  string  `json:"employee_id"`
		JobID       string  `json:"job_id"`
		FirstName   string  `json:"first_name"`
		LastName    string  `json:"last_name"`
		Email       string  `json:"email"`
		PhoneNumber string  `json:"phone_number"`
		Salary      float64 `json:"salary"`
		entity.TimeStamp
	}

	EmployeePaginationResponse struct {
		Data []AllEmployeeResponse `json:"data"`
		PaginationResponse
	}

	GetAllEmployeeRepositoryResponse struct {
		Employees []entity.Employee
		PaginationResponse
	}

	UpdateProfileRequest struct {
		EmployeeID  string  `json:"employee_id"`
		JobID       string  `json:"job_id"`
		FirstName   string  `json:"first_name,omitempty"`
		LastName    string  `json:"last_name,omitempty"`
		Email       string  `json:"email,omitempty"`
		PhoneNumber string  `json:"phone_number,omitempty"`
		Salary      float64 `json:"salary,omitempty"`
	}

	DeleteEmployeeRequest struct {
		EmployeeID string `json:"employee_id"`
	}
)
