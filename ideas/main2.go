package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type GraphQL struct {
	Schema  *graphql.Schema
	Logger  logger
}

// nolint: gas
func main() {
	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	page, err := getFile("./graphql.html")
	schema, err := getFile("./schema.graphql")
	if err != nil {
		panic(err)
	}

	// Create the request handler; inject dependencies.
	h := handler.GraphQL{
		// Parse and validate schema. Panic if unable to do so.
		Schema:  graphql.MustParseSchema(schema.String(), root),
		
	}

	// Register handlers to routes.
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	// Configure the HTTP server.
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}

func getFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func (h GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validate the request.
	if ok := isSupported(r.Method); !ok {
		respond(w, errorJSON("only POST or GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}

	req, err := parse(r)
	if err != nil {
		respond(w, errorJSON(err.Error()), http.StatusBadRequest)
		return
	}

	n := len(req.queries)
	if n == 0 {
		respond(w, errorJSON("no queries to execute"), http.StatusBadRequest)
		return
	}

	// NOTE: User authentication should happen here, if needed.
	//
	// Authentication determines who the request originated from.
	// Authorization business logic (in services, not this API) will rely on this authentication data.
	//
	// We don't need it for this example application, but a typical production application would
	// perform request authentication.
	//
	// The result of authentication should probably be placed on the request context so it can be
	// passed to resolvers and loaders.

	// Here, begin request execution...
	var (
		ctx       = h.Loaders.Attach(r.Context()) // Attach dataloaders onto the request context.
		responses = make([]*graphql.Response, n)  // Allocate a slice large enough for all responses.
		wg        sync.WaitGroup                  // Use the WaitGroup to wait for all executions to finish.
	)

	wg.Add(n)

	for i, q := range req.queries {
		// Loop through the parsed queries from the request.
		// These queries are executed in separate goroutines so they process in parallel.
		go func(i int, q query) {
			res := h.Schema.Exec(ctx, q.Query, q.OpName, q.Variables)

			// We have to do some work here to expand errors when it is possible for a resolver to return
			// more than one error (for example, a list resolver).
			res.Errors = errors.Expand(res.Errors)

			responses[i] = res
			wg.Done()
		}(i, q)
	}

	wg.Wait()

	// NOTE: Before returning the response to the user, we can inspect the results for errors
	// and log them.
	//
	// However, be mindful that the standard go log package uses a global mutex to protect writes
	// to stdout. In a log-happy service, you may see service goroutines start to block on that mutex.

	// After we've doctored up our response by filtering internal error messages or adding data to
	// the 'extensions' field, we marshal the response to JSON.
	var resp []byte
	if req.isBatch {
		resp, err = json.Marshal(responses)
	} else if len(responses) > 0 {
		resp, err = json.Marshal(responses[0])
	}

	if err != nil {
		respond(w, errorJSON("server error"), http.StatusInternalServerError)
		return
	}

	respond(w, resp, http.StatusOK)
}

// logger defines an interface with a single method.
type logger interface {
	Printf(fmt string, values ...interface{})
}

// A request respresents an HTTP request to the GraphQL endpoint.
// A request can have a single query or a batch of requests with one or more queries.
// It is important to distinguish between a single query request and a batch request with a single query.
// The shape of the response will differ in both cases.
type request struct {
	queries []query
	isBatch bool
}

// A query represents a single GraphQL query.
type query struct {
	OpName    string                 `json:"operationName"`
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
