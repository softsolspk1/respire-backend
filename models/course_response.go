package models

// CourseResponse is used for API responses for a single course
type CourseResponse struct {
	Code   int    `json:"code"`
	Error  string `json:"error,omitempty"`
	Result Course `json:"result,omitempty"`
}