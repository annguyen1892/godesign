// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "reviews": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/annguyen1892/godesign/design
// --out=$(GOPATH)/src/reviews
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// review media type (default view)
//
// Identifier: application/vnd.gophercon.goa.review; view=default
type Review struct {
	// Content of review
	Content string `form:"content" json:"content" xml:"content"`
	// customer_id create review
	CustomerID int `form:"customer_id" json:"customer_id" xml:"customer_id"`
	// product_id of review
	ProductID int `form:"product_id" json:"product_id" xml:"product_id"`
	// rating create review
	Rating int `form:"rating" json:"rating" xml:"rating"`
	// Title of review
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the Review media type instance.
func (mt *Review) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Content == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "content"))
	}

	if utf8.RuneCountInString(mt.Content) < 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.content`, mt.Content, utf8.RuneCountInString(mt.Content), 50, true))
	}
	if mt.CustomerID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.customer_id`, mt.CustomerID, 1, true))
	}
	if mt.ProductID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.product_id`, mt.ProductID, 1, true))
	}
	if mt.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 1, true))
	}
	if mt.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 5, false))
	}
	if utf8.RuneCountInString(mt.Title) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 5, true))
	}
	return
}

// DecodeReview decodes the Review instance encoded in resp body.
func (c *Client) DecodeReview(resp *http.Response) (*Review, error) {
	var decoded Review
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
