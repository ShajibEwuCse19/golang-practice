package product

// Handler handles HTTP requests for products and depends on the Service contract.
type Handler struct {
	svc Service // Local Service interface contract defined in port.go
}

// NewHandler initializes a new Handler with the provided Service implementation.
func NewHandler(svc Service) *Handler { 
	return &Handler{
		svc: svc,
	}
}