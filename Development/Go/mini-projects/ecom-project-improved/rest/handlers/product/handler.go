package product

type Handler struct {
	svc Service // own contract -> rest/handlers/product/port.go
}

// During constructor create in serve.go, using same signature as own contract which is services/service.go
// services/service.go has implementation which produce data for handler/port.go contract list. 
func NewHandler(svc Service) *Handler { 
	return &Handler{
		svc: svc,
	}
}