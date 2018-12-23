package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/vlamug/elementary-cacher/api"
	"github.com/vlamug/elementary-cacher/cache"
	"github.com/vlamug/elementary-cacher/pkg/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultListenAddr = ":9071"
	defaultCacheSize  = 100000
)

func main() {
	var (
		listenAddr string
		cacheSize  int
	)

	// init flags
	flag.StringVar(&listenAddr, "listen.addr", defaultListenAddr, "Address to listen requests")
	flag.IntVar(&cacheSize, "cache.size", defaultCacheSize, "Cache size")
	flag.Parse()

	// init cacher
	cacher := cache.NewMem(cacheSize)

	// init handlers
	http.HandleFunc("/get", middleware.MetricsMiddleware("get", api.GetHandler(cacher)))
	http.HandleFunc("/set", middleware.MetricsMiddleware("set", api.SetHandler(cacher)))

	// init metrics handler
	http.Handle("/metrics", promhttp.Handler())

	// run server
	log.Printf("running on address: %s\n", listenAddr)
	log.Printf("cache size: %d\n", cacheSize)
	log.Fatalln(http.ListenAndServe(listenAddr, nil))
}
