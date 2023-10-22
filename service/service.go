package service

type Service struct {
	// The name of the service
	Name string

	// The service type
	ServiceType ServiceTypes
}

// new
func New(name string, serviceType ServiceTypes) *Service {
	return &Service{
		Name:        name,
		ServiceType: serviceType,
	}
}
