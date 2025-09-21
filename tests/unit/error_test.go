package unit_test

import (
	"net/http"

	apperrors "github.com/STommydx/VibeFolio/src/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Handling", func() {
	Describe("AppError", func() {
		Context("when creating a new AppError", func() {
			It("should create an error with the correct properties", func() {
				err := apperrors.NewAppError(400, "Bad Request", "Invalid input")

				Expect(err.Code).To(Equal(400))
				Expect(err.Message).To(Equal("Bad Request"))
				Expect(err.Details).To(Equal("Invalid input"))
				Expect(err.Error()).To(Equal("Bad Request: Invalid input"))
			})

			It("should create an error without details", func() {
				err := apperrors.NewAppError(500, "Internal Server Error", "")

				Expect(err.Code).To(Equal(500))
				Expect(err.Message).To(Equal("Internal Server Error"))
				Expect(err.Details).To(BeEmpty())
				Expect(err.Error()).To(Equal("Internal Server Error"))
			})
		})

		Context("when creating specific error types", func() {
			It("should create a BadRequestError", func() {
				err := apperrors.NewBadRequestError("Invalid request", "Missing required field")

				Expect(err.Code).To(Equal(http.StatusBadRequest))
				Expect(err.Message).To(Equal("Invalid request"))
				Expect(err.Details).To(Equal("Missing required field"))
			})

			It("should create a NotFoundError", func() {
				err := apperrors.NewNotFoundError("Resource not found", "User with ID 123 not found")

				Expect(err.Code).To(Equal(http.StatusNotFound))
				Expect(err.Message).To(Equal("Resource not found"))
				Expect(err.Details).To(Equal("User with ID 123 not found"))
			})

			It("should create an InternalServerError", func() {
				err := apperrors.NewInternalServerError("Internal error", "Database connection failed")

				Expect(err.Code).To(Equal(http.StatusInternalServerError))
				Expect(err.Message).To(Equal("Internal error"))
				Expect(err.Details).To(Equal("Database connection failed"))
			})

			It("should create a ServiceUnavailableError", func() {
				err := apperrors.NewServiceUnavailableError("Service unavailable", "Database is down")

				Expect(err.Code).To(Equal(http.StatusServiceUnavailable))
				Expect(err.Message).To(Equal("Service unavailable"))
				Expect(err.Details).To(Equal("Database is down"))
			})
		})
	})

	Describe("HTTP Status Code", func() {
		It("should return the correct HTTP status code", func() {
			err := apperrors.NewAppError(418, "I'm a teapot", "")

			Expect(err.HTTPStatusCode()).To(Equal(418))
		})
	})
})
