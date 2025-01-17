package handler

import (
	"net/http"

	"github.com/CarlosGenuino/uber-go/internal/service"
	"github.com/gin-gonic/gin"
)

type PassengerHandler struct {
	passengerService *service.PassengerService
}

func NewPassengerHandler(passengerService *service.PassengerService) *PassengerHandler {
	return &PassengerHandler{passengerService: passengerService}
}

func (h *PassengerHandler) CreatePassenger(c *gin.Context) {
	var req struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passenger, err := h.passengerService.CreatePassenger(req.ID, req.Name, req.Latitude, req.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, passenger)
}

func (h *PassengerHandler) GetPassenger(c *gin.Context) {
	id := c.Param("id")

	passenger, err := h.passengerService.GetPassenger(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passenger)
}
