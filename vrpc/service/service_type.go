package service

type ServiceTypes struct {
	// The name of the service type
	Name string
}

var (
	// a list of all known service types
	serviceTypes = make(map[string]ServiceTypes)
)

// RegisterServiceType registers a new service type
func RegisterServiceType(name string) error {
	if _, ok := serviceTypes[name]; ok {
		return nil
	}
	serviceTypes[name] = ServiceTypes{Name: name}
	return nil
}
