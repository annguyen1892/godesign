//go:generate goagen bootstrap -d github.com/annguyen1892/godesign/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"reviews/app"
)

func main() {
	// Create service
	service := goa.New("reviews")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "review" controller
	c := NewReviewController(service)
	app.MountReviewController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
