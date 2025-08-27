package errors

import "fmt"

// Error returns the error message for the bin not found error.
func (e *BinNotFoundError) Error() string {
	return fmt.Sprintf("bin not found: %s", e.BinID)
}

// Error returns the error message for the file not found error.
func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s/%s", e.BinID, e.Filename)
}


// Error returns the error message for the API error.
func (e *APIError) Error() string {
	return fmt.Sprintf("API error (%d): %s", e.StatusCode, e.Message)
}