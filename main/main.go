package main

import (
	"github.com/CarlosGenuino/uber-go/internal/handler"
	"github.com/CarlosGenuino/uber-go/internal/repository"
	"github.com/CarlosGenuino/uber-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the repository
	passengerRepo := repository.NewPassengerRepository()
	driverRepo := repository.NewDriverRepository()
	rideRepo := repository.NewRideRepository()

	// Initialize the service
	passengerService := service.NewPassengerService(passengerRepo)
	driverService := service.NewDriverService(driverRepo)
	rideService := service.NewRideService(passengerService, driverService, rideRepo)
	// Initialize the handler
	passengerHandler := handler.NewPassengerHandler(passengerService)
	driverHandler := handler.NewDriverHandler(driverService)
	rideHandler := handler.NewRideHandler(rideService)

	// Set up the Gin router
	r := gin.Default()

	// Register the routes
	r.POST("/passengers", passengerHandler.CreatePassenger)
	r.GET("/passengers/:id", passengerHandler.GetPassenger)

	r.POST("/drivers", driverHandler.CreateDriver)
	r.GET("/drivers/:id", driverHandler.GetDriver)

	r.POST("/rides", rideHandler.RequestRide)
	r.POST("/rides/end", rideHandler.EndRide)
	r.POST("/rides/cancel", rideHandler.CancelRide)

	// Start the server
	r.Run(":8080")
}
