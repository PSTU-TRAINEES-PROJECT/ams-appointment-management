package controller

import (
	"ams-appointment-management/app/service"
	"net/http"

	"ams-appointment-management/app/common/logger"
	"ams-appointment-management/app/common/pagination"
	"ams-appointment-management/app/common/utils"
	"ams-appointment-management/app/controller/response"
	"ams-appointment-management/app/serializer"

	"github.com/labstack/echo/v4"
)

type AppointmentControllerInterface interface {
	CreateUser(context echo.Context) error
	FindAll(context echo.Context) error
	FindByID(context echo.Context) error
	Update(context echo.Context) error
	Delete(context echo.Context) error
}

type AppointmentController struct {
	appointmentService service.AppointmentServiceInterface
}

func NewAppointmentController(appointmentService service.AppointmentServiceInterface) AppointmentController {
	return AppointmentController{appointmentService: appointmentService}
}

func (c *AppointmentController) Create(context echo.Context) error {
	newAppointment := serializer.Appointment{}
	if err := context.Bind(&newAppointment); err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(response.ErrParsingRequestBody))
	}
	user, err := c.appointmentService.Create(&newAppointment)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}
	return context.JSON(http.StatusCreated, response.GenerateSuccessResponse("created successfully", user.Id))
}

func (c *AppointmentController) FindAll(context echo.Context) error {
	page := pagination.Page{}
	pageinfo, err := page.GetPageInformation(context)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}
	pageResp := pagination.PageResponse{}
	pageResp.Data, pageResp.Count, err = c.appointmentService.FindAll(*pageinfo)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}
	return context.JSON(http.StatusOK, response.GenerateSuccessResponse("successful", pageResp))
}

func (c *AppointmentController) FindByID(context echo.Context) error {
	ID, err := utils.ParseQueryParamAsInt(context, "id")
	//value, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(response.ErrInvalidRequestParams))
	}

	user, err := c.appointmentService.FindByID(ID)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}
	return context.JSON(http.StatusOK, response.GenerateSuccessResponse("successful", user))
}

func (c *AppointmentController) Update(context echo.Context) error {
	ID, err := utils.ParseQueryParamAsInt(context, "id")
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(response.ErrInvalidRequestParams))
	}
	updatedUser := serializer.Appointment{}
	if err := context.Bind(&updatedUser); err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(response.ErrParsingRequestBody))
	}
	user, err := c.appointmentService.Update(ID, &updatedUser)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}
	return context.JSON(http.StatusAccepted, response.GenerateSuccessResponse("UserUpdatedSuccessfully", user.Id))
}

func (c *AppointmentController) Delete(context echo.Context) error {
	ID, err := utils.ParseQueryParamAsInt(context, "id")
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(response.ErrInvalidRequestParams))
	}
	ID, err = c.appointmentService.Delete(ID)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}
	return context.JSON(http.StatusAccepted, response.GenerateSuccessResponse("UserDeletedSuccessfully", ID))
}
