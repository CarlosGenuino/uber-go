package main

import (
	"log"

	"github.com/CarlosGenuino/uber-go/internal/config"
	"github.com/CarlosGenuino/uber-go/internal/handler"
	"github.com/CarlosGenuino/uber-go/internal/repository"
	"github.com/CarlosGenuino/uber-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load database configuration
	dbConfig := config.LoadDBConfig()

	// Connect to the database
	db, err := sqlx.Connect("postgres", dbConfig.GetConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize the repositories
	passengerRepo := repository.NewPassengerRepository(db)
	driverRepo := repository.NewDriverRepository(db)
	rideRepo := repository.NewRideRepository(db)

	// Initialize the services
	passengerService := service.NewPassengerService(passengerRepo)
	driverService := service.NewDriverService(driverRepo)
	rideService := service.NewRideService(passengerService, driverService, rideRepo)

	// Initialize the handlers
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
	r.POST("/rides/accept", rideHandler.AcceptRide)
	r.POST("/rides/end", rideHandler.EndRide)
	r.POST("/rides/cancel", rideHandler.CancelRide)

	// Start the server
	r.Run(":8080")
}
