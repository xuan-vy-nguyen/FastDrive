package models

// ImageRequest is ok
type ImageRequest struct {
	Name string `json:"name"`
}

// ImageDB is ok
type ImageDB struct {
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Image []byte `json:"image"`
}
