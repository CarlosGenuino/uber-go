package handler

import (
	"net/http"

	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/internal/service"
	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	driverService *service.DriverService
}

func NewDriverHandler(driverService *service.DriverService) *DriverHandler {
	return &DriverHandler{driverService: driverService}
}

func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var req struct {
		Name      string  `json:"name"`
		LicenseID string  `json:"license_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Car       struct {
			Make  string `json:"make"`
			Model string `json:"model"`
			Year  int    `json:"year"`
		} `json:"car"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := domain.NewCar(req.Car.Make, req.Car.Model, req.Car.Year)
	if !domain.ValidateCar(car) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car details"})
		return
	}

	driver, err := h.driverService.CreateDriver(req.Name, req.LicenseID, req.Latitude, req.Longitude, car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, driver)
}

func (h *DriverHandler) GetDriver(c *gin.Context) {
	id := c.Param("id")

	driver, err := h.driverService.GetDriver(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}
