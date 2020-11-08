package pipe_filter

// Request data interface
type Request interface {}

// Response data interface
type Response interface{}

// Filter interface
type Filter interface {
	// GetName get name of filter.
	GetName() string

	// Process process request and return response.
	Process(data Request) (Response, error)
}