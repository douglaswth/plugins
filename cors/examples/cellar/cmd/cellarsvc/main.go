package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware/debugging"
	"goa.design/goa/http/middleware/logging"
	cellar "goa.design/plugins/cors/examples/cellar"
	sommeliersvr "goa.design/plugins/cors/examples/cellar/gen/http/sommelier/server"
	storagesvr "goa.design/plugins/cors/examples/cellar/gen/http/storage/server"
	swaggersvr "goa.design/plugins/cors/examples/cellar/gen/http/swagger/server"
	sommelier "goa.design/plugins/cors/examples/cellar/gen/sommelier"
	storage "goa.design/plugins/cors/examples/cellar/gen/storage"
	swagger "goa.design/plugins/cors/examples/cellar/gen/swagger"
)

func main() {
	// Define command line flags, add any other flag required to configure
	// the service.
	var (
		addr = flag.String("listen", ":8080", "HTTP listen `address`")
		dbg  = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice. The goa.design/middleware/logging/...
	// packages define log adapters for common log packages.
	var (
		logger  *log.Logger
		adapter logging.Adapter
	)
	{
		logger = log.New(os.Stderr, "[cellar] ", log.Ltime)
		adapter = logging.Adapt(logger)
	}

	// Create the structs that implement the services.
	var (
		sommelierSvc sommelier.Service
		storageSvc   storage.Service
		swaggerSvc   swagger.Service
	)
	{
		sommelierSvc = cellar.NewSommelier(logger)
		storageSvc = cellar.NewStorage(logger)
		swaggerSvc = cellar.NewSwagger(logger)
	}

	// Wrap the services in endpoints that can be invoked from other
	// services potentially running in different processes.
	var (
		sommelierEndpoints *sommelier.Endpoints
		storageEndpoints   *storage.Endpoints
		swaggerEndpoints   *swagger.Endpoints
	)
	{
		sommelierEndpoints = sommelier.NewEndpoints(sommelierSvc)
		storageEndpoints = storage.NewEndpoints(storageSvc)
		swaggerEndpoints = swagger.NewEndpoints(swaggerSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		sommelierServer *sommeliersvr.Server
		storageServer   *storagesvr.Server
		swaggerServer   *swaggersvr.Server
	)
	{
		sommelierServer = sommeliersvr.New(sommelierEndpoints, mux, dec, enc)
		storageServer = storagesvr.New(storageEndpoints, mux, dec, enc)
		swaggerServer = swaggersvr.New(swaggerEndpoints, mux, dec, enc)
	}

	// Configure the mux.
	sommeliersvr.Mount(mux, sommelierServer)
	storagesvr.Mount(mux, storageServer)
	swaggersvr.Mount(mux, swaggerServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if *dbg {
			handler = debugging.New(mux, adapter)(handler)
		}
		handler = logging.New(adapter)(handler)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the service to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: *addr, Handler: handler}
	go func() {
		for _, m := range sommelierServer.Mounts {
			logger.Printf("[INFO] service %q method %q mounted on %s %s", sommelierServer.Service(), m.Method, m.Verb, m.Pattern)
		}
		for _, m := range storageServer.Mounts {
			logger.Printf("[INFO] service %q method %q mounted on %s %s", storageServer.Service(), m.Method, m.Verb, m.Pattern)
		}
		for _, m := range swaggerServer.Mounts {
			logger.Printf("[INFO] service %q file %q mounted on %s %s", swaggerServer.Service(), m.Method, m.Verb, m.Pattern)
		}
		logger.Printf("[INFO] listening on %s", *addr)
		errc <- srv.ListenAndServe()
	}()

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Shutdown gracefully with a 30s timeout.
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)

	logger.Println("exited")
}
