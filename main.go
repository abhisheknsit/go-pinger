package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"

	"github.com/abhisheknsit/go-pinger/metrics"
)

var (
	addr_metrics = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")
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
	log.Fatal(http.ListenAndServe(*addr_metrics, nil))
}
