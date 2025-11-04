package ginrouter

import (
	"net/http"
	"strconv"

	"github.com/anggaaryas/go-mockapi"
	"github.com/gin-gonic/gin"
)

type config struct {
	r *gin.Engine
}

func Create(r *gin.Engine) mockapi.Router {
	return &config{
		r: r,
	}
}

func (cfg *config) getErrorResponse(err error) mockapi.APIError {
	statusCode := 500

	if customErr, ok := err.(CustomError); ok {
		statusCode = customErr.StatusCode()
		return mockapi.APIError{
			StatusCode: statusCode,
			Message:    customErr.Error(),
		}
	}

	if gin.Mode() != gin.ReleaseMode {
		return mockapi.APIError{
			StatusCode: statusCode,
			Message:    err.Error(),
		}
	}
	return mockapi.APIError{
		StatusCode: statusCode,
		Message:    "An error occurred while processing your request",
	}
}

func (cfg *config) SetupMockApiRoute(service mockapi.Service) error {

	cfg.r.StaticFS("/static", http.FS(mockapi.GetStaticFiles()))

	api := cfg.r.Group("/api")

	api.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			apiErr := cfg.getErrorResponse(NewIDShouldBeIntError("id"))
			c.JSON(apiErr.StatusCode, apiErr)
			return
		}
		book, err := service.GetBookByID(id)
		if err != nil {
			apiErr := cfg.getErrorResponse(err)
			c.JSON(apiErr.StatusCode, apiErr)
			return
		}
		c.JSON(200, book)
	})
	api.GET("/books", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
		search := c.DefaultQuery("search", "")
		books, err := service.GetBooks(page, pageSize, search)
		if err != nil {
			apiErr := cfg.getErrorResponse(err)
			c.JSON(apiErr.StatusCode, apiErr)
			return
		}
		c.JSON(200, books)
	})

	return nil
}
