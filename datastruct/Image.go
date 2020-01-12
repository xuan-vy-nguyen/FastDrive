package datastruct

// ImageRequest is ok
type ImageRequest struct {
	Name string `json:"name"`
}

// ImageDB is ok
type ImageDB struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image []byte `json:"image"`
}
