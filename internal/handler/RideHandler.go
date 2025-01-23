package handler

import (
	"net/http"

	"github.com/CarlosGenuino/uber-go/internal/service"
	"github.com/gin-gonic/gin"
)

type RideHandler struct {
	rideService *service.RideService
}

func NewRideHandler(rideService *service.RideService) *RideHandler {
	return &RideHandler{rideService: rideService}
}

func (h *RideHandler) RequestRide(c *gin.Context) {
	var req struct {
		PassengerID string `json:"passenger_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ride, err := h.rideService.RequestRide(req.PassengerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ride)
}

func (h *RideHandler) AcceptRide(c *gin.Context) {
	var req struct {
		RideID   string `json:"ride_id"`
		DriverID string `json:"driver_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ride, err := h.rideService.AcceptRide(req.RideID, req.DriverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ride)
}

func (h *RideHandler) EndRide(c *gin.Context) {
	var req struct {
		RideID string `json:"ride_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ride, err := h.rideService.EndRide(req.RideID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ride)
}

func (h *RideHandler) CancelRide(c *gin.Context) {
	var req struct {
		RideID string `json:"ride_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ride, err := h.rideService.CancelRide(req.RideID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ride)
}
