package main

import (
	"flight-management/booking/dao"
	"flight-management/booking/handler"
	"flight-management/common/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Setup()
	config.GetDatabaseConnection()

	dbConfig := config.GetDatabaseConnection()
	bookingService := dao.NewBookingService(dbConfig)

	r := gin.Default()
	h := handler.NewHandler(&bookingService)
	r.GET("/ping", h.HandleHealthCheck)

	r.POST("/booking", h.HandleReserveBooking)
	r.GET("/booking", h.HandlerGetListBooking)
	r.GET("/booking/:id", h.HandleGetBookingById)
	err := r.Run("0.0.0.0:8082")
	if err != nil {
		panic(err)
	}
}
