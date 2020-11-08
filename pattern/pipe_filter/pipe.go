package pipe_filter

// Pipe interface
type Pipe interface {
	// GetName get name of pipe.
	GetName() string

	// Run start run filters.
	Run(filters ...Filter) (Response, error)
}