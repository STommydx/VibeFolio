package unit_test

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load Testing", func() {
	Describe("Concurrent Health Checks", func() {
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

		Context("when handling high concurrent load", func() {
			It("should handle 1000 concurrent requests with < 5% error rate", func() {
				client := &http.Client{
					Timeout: 5 * time.Second,
				}

				const numRequests = 1000
				const maxErrorRate = 0.05 // 5%

				var wg sync.WaitGroup
				results := make(chan bool, numRequests)

				// Launch 1000 concurrent requests
				start := time.Now()
				for i := 0; i < numRequests; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						resp, err := client.Get(server.URL + "/healthz")
						if err == nil && resp.StatusCode == http.StatusOK {
							results <- true
						} else {
							results <- false
						}
					}()
				}

				// Close results channel when all requests are done
				go func() {
					wg.Wait()
					close(results)
				}()

				// Collect results
				successCount := 0
				totalCount := 0
				for result := range results {
					totalCount++
					if result {
						successCount++
					}
				}
				end := time.Now()

				// Calculate error rate
				errorRate := float64(totalCount-successCount) / float64(totalCount)

				// Assertions
				Expect(totalCount).To(Equal(numRequests))
				Expect(errorRate).To(BeNumerically("<=", maxErrorRate))
				Expect(end.Sub(start)).To(BeNumerically("<", 10*time.Second))
			})
		})

		Context("when handling sustained load", func() {
			It("should maintain consistent response times over 30 seconds", func() {
				client := &http.Client{
					Timeout: 5 * time.Second,
				}

				const duration = 30 * time.Second
				const interval = 100 * time.Millisecond // 10 requests per second
				const maxAvgResponseTime = 50 * time.Millisecond

				var responseTimes []time.Duration
				done := make(chan bool)

				// Start making requests for 30 seconds
				ticker := time.NewTicker(interval)
				defer ticker.Stop()

				go func() {
					for {
						select {
						case <-ticker.C:
							reqStart := time.Now()
							resp, err := client.Get(server.URL + "/healthz")
							reqEnd := time.Now()

							if err == nil && resp.StatusCode == http.StatusOK {
								responseTimes = append(responseTimes, reqEnd.Sub(reqStart))
							}
						case <-done:
							return
						}
					}
				}()

				// Wait for 30 seconds
				time.Sleep(duration)
				close(done)

				// Calculate average response time
				if len(responseTimes) > 0 {
					var total time.Duration
					for _, rt := range responseTimes {
						total += rt
					}
					avgResponseTime := total / time.Duration(len(responseTimes))

					Expect(avgResponseTime).To(BeNumerically("<=", maxAvgResponseTime))
				}
			})
		})
	})
})
