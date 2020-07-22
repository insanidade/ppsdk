package interfaces

type ResponseBodyRoot interface {
	// GetId() string
	GetLinks() map[string]Link
	GetDebugID() string
}
