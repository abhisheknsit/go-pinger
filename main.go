package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"

	"github.com/abhisheknsit/go-pinger/metrics"
)

var (
	addr              = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")
	uniformDomain     = flag.Float64("uniform.domain", 0.0002, "The domain for the uniform distribution.")
	normDomain        = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean          = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
)

func main() {
	flag.Parse()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		timer := prometheus.NewTimer(metrics.HandlerTimer.WithLabelValues("ping"))
		defer timer.ObserveDuration()
		c.JSON(200, gin.H{
			"message": "pong",
		})
		metrics.RequestCounter.WithLabelValues("ping").Inc()
	})

	go r.Run()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
