package contract_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/STommydx/VibeFolio/src/controllers"
	"github.com/STommydx/VibeFolio/src/services"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Healthz Endpoint Contract", func() {
	var (
		server   *httptest.Server
		client   *http.Client
		request  *http.Request
		response *http.Response
	)

	BeforeEach(func() {
		// Set up the actual server with our implementation
		mux := http.NewServeMux()
		api := humago.New(mux, huma.DefaultConfig("VibeFolio Health Check", "1.0.0"))

		// Create services and controllers
		healthService := services.NewHealthService("1.0.0")
		healthController := controllers.NewHealthController(healthService)

		// Register routes
		healthController.RegisterRoutes(api)

		server = httptest.NewServer(mux)
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	})

	AfterEach(func() {
		if server != nil {
			server.Close()
		}
		if response != nil {
			response.Body.Close()
		}
	})

	Describe("GET /healthz", func() {
		Context("when the service is healthy", func() {
			BeforeEach(func() {
				var err error
				request, err = http.NewRequest("GET", server.URL+"/healthz", nil)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns 200 status code", func() {
				var err error
				response, err = client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns Content-Type application/json", func() {
				var err error
				response, err = client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.Header.Get("Content-Type")).To(ContainSubstring("application/json"))
			})

			It("returns valid HealthStatus JSON structure", func() {
				var err error
				response, err = client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
