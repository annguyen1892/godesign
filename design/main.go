
package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("cellar", func() {
	Description("The product review service")
	Host("localhost:8080")
})

var productReviewPayload = Type("productReviewPayload", func() {
	Description("ProductReviewPayload is the type used")

	Attribute("title", String, "Title of review", func() {
		MinLength(5)
	})

	Attribute("content", String, "Content of review", func() {
		MinLength(50)
	})

	Attribute("product_id", Integer, "product_id of review", func() {
		Minimum(1)
	})

	Attribute("customer_id", Integer, "customer_id create review", func() {
		Minimum(1)
	})

	Attribute("rating", Integer, "rating create review", func() {
		Minimum(1)
		Maximun(5)
	})

	Required("title", "content", "rating", "customer_id")
})

var reviewMedia = MediaType("application/vnd.gophercon.goa.review", func() {
	TypeName("review")
	Reference(productReviewPayload)

	Attributes(func() {
		Attribute("title")
		Attribute("content")
		Attribute("product_id")
		Attribute("customer_id")
		Attribute("rating")
		Required("title", "content", "product_id", "customer_id", "rating")
	})
})

var _ = Resource("review", func() {
	Description("review")
	BasePath("/reviews")

	Action("create", func() {
		Description("creates a review")
		Routing(POST("/"))
		Payload(productReviewPayload)
		Response(Created)
	})

	Action("show", func() {
		Description("Get a review")
		Routing(GET("/:id"))
    Params(function(){
      Param("id", Integer, "Review Id")
    })
		Response(OK, productReviewPayload)
	})
})
