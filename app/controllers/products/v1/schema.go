package products

// Generic schema's
type ProductRequest struct {
	Name      string                 `json:"name"`
	Metadadta map[string]interface{} `json:"metadata"`
}

type ProductResponse struct {
	UID       string                 `json:"uid"`
	Slug      string                 `json:"slug"`
	Name      string                 `json:"name"`
	Metadadta map[string]interface{} `json:"metadata"`
}

// Create product request schema
type CreateProductRequest struct {
	Slug string `json:"slug"`
	ProductRequest
}

// Update product request schema
type UpdateProductRequest struct {
	ProductRequest
}
