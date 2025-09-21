package unit_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Performance Tests", func() {
	Describe("Health Endpoint Performance", func() {
		var server *httptest.Server

		BeforeEach(func() {
			// Create a test server with a simple health endpoint
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/healthz" {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"status":"healthy","timestamp":"2025-09-21T10:00:00Z","version":"1.0.0"}`))
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			}))
		})

		AfterEach(func() {
			server.Close()
		})

		Context("when measuring response time", func() {
			It("should respond within 100ms", func() {
				client := &http.Client{
					Timeout: 5 * time.Second,
				}

				start := time.Now()
				resp, err := client.Get(server.URL + "/healthz")
				end := time.Now()

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				Expect(end.Sub(start)).To(BeNumerically("<", 100*time.Millisecond))
			})
		})

		Context("when making multiple concurrent requests", func() {
			It("should handle 100 concurrent requests within 1 second", func() {
				client := &http.Client{
					Timeout: 5 * time.Second,
				}

				// Make 100 concurrent requests
				done := make(chan bool, 100)
				start := time.Now()

				for i := 0; i < 100; i++ {
					go func() {
						resp, err := client.Get(server.URL + "/healthz")
						if err == nil && resp.StatusCode == http.StatusOK {
							done <- true
						} else {
							done <- false
						}
					}()
				}

				// Wait for all requests to complete
				successCount := 0
				for i := 0; i < 100; i++ {
					if <-done {
						successCount++
					}
				}

				end := time.Now()

				// All requests should succeed
				Expect(successCount).To(Equal(100))
				// All requests should complete within 1 second
				Expect(end.Sub(start)).To(BeNumerically("<", 1*time.Second))
			})
		})
	})
})
