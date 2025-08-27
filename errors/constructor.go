// Package errors provides error types for the filebin application.
package errors

// NewBinNotFoundError creates a new BinNotFoundError for a bin not found error.
func NewBinNotFoundError(binID string) *BinNotFoundError {
	return &BinNotFoundError{
		BinID: binID,
	}
}

// NewFileNotFoundError creates a new FileNotFoundError for a file not found error.
func NewFileNotFoundError(binID, filename string) *FileNotFoundError {
	return &FileNotFoundError{
		BinID:    binID,
		Filename: filename,
	}
}

// NewAPIError creates a new APIError for an API error with the given status code and message.
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}
