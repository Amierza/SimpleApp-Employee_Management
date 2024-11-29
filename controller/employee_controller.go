package controller

import (
	"net/http"

	"github.com/Amierza/employee-management/dto"
	"github.com/Amierza/employee-management/service"
	"github.com/Amierza/employee-management/utils"
	"github.com/gin-gonic/gin"
)

type (
	EmployeeController interface {
		CreateEmployee(*gin.Context)
		GetAllEmployee(*gin.Context)
		UpdateProfile(*gin.Context)
		DeleteEmployee(*gin.Context)
	}

	employeeController struct {
		employeeService service.EmployeeService
	}
)

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &employeeController{
		employeeService: employeeService,
	}
}

func (c *employeeController) CreateEmployee(ctx *gin.Context) {
	var employee dto.EmployeeCreateRequest
	if err := ctx.ShouldBind(&employee); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.CreateEmployee(ctx.Request.Context(), employee)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_EMPLOYEE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_EMPLOYEE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) GetAllEmployee(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.GetAllEmployeeWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_EMPLOYEE,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *employeeController) UpdateProfile(ctx *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.UpdateProfileEmployee(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_PROFILE_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_PROFILE_EMPLOYEE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) DeleteEmployee(ctx *gin.Context) {
	var req dto.DeleteEmployeeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.employeeService.DeleteEmployee(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_EMPLOYEE, nil)
	ctx.JSON(http.StatusOK, res)
}
