package product

// Response represents a generic response format for API responses.
type Response struct {
	Status  int         `json:"status"`  // HTTP status code of the response
	Message string      `json:"message"` // Message describing the result of the operation
	Data    interface{} `json:"data"`    // Data payload of the response
}
