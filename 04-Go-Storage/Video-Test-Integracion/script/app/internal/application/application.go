package application

// Application is the application interface.
type Application interface {
	// SetUp sets up the application.
	SetUp() (err error)
	// Run runs the application.
	Run() (err error)
}