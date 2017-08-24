package main

import (
	"app"
	"github.com/goadesign/goa"
)

// ReviewController implements the review resource.
type ReviewController struct {
	*goa.Controller
}

// NewReviewController creates a review controller.
func NewReviewController(service *goa.Service) *ReviewController {
	return &ReviewController{Controller: service.NewController("ReviewController")}
}

// Create runs the create action.
func (c *ReviewController) Create(ctx *app.CreateReviewContext) error {
	// ReviewController_Create: start_implement

	// Put your logic here

	// ReviewController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *ReviewController) Show(ctx *app.ShowReviewContext) error {
	// ReviewController_Show: start_implement

	// Put your logic here

	// ReviewController_Show: end_implement
	return nil
}
