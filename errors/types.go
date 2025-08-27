package errors

// FilebinError represents a generic filebin error.
type FilebinError struct {
	Message string
	Code    int
}

// BinNotFoundError represents a bin not found error.
type BinNotFoundError struct {
	BinID string
}

// FileNotFoundError represents a file not found error.
type FileNotFoundError struct {
	BinID    string
	Filename string
}


// APIError represents an API error.
type APIError struct {
	StatusCode int
	Message    string
}