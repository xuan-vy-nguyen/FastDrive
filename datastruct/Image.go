package datastruct

// ImageRequest is ok
type ImageRequest struct {
	Name string `json:"name"`
}

// ImageResponse is ok
type ImageResponse struct {
	Length int      `json:"length"`
	Names  []string `json:"names"`
	Links  []string `json:"links"`
}
