package integration_test

import (
	"net/http"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Check Integration", func() {
	var (
		client   *http.Client
		response *http.Response
	)

	BeforeEach(func() {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	})

	AfterEach(func() {
		if response != nil {
			response.Body.Close()
		}
	})

	Describe("Health check endpoint", func() {
		Context("when the server is running", func() {
			It("should respond to GET /healthz requests", func() {
				// Skip this test if we're not running integration tests
				if os.Getenv("INTEGRATION_TESTS") != "true" {
					Skip("Skipping integration test. Set INTEGRATION_TESTS=true to run.")
				}

				// This test should pass now that we have implemented the server
				// We'll use localhost:9090 as the default port
				var err error
				response, err = client.Get("http://localhost:9090/healthz")

				// This should succeed because the server is running
				Expect(err).NotTo(HaveOccurred())
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
