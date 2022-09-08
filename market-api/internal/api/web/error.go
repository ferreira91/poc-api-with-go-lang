package web

var (
	ErrBadRequest          = "invalid_request"
	ErrNotFound            = "not_found"
	ErrInternalServerError = "server_error"
	BadRequest             = "Invalid request"
	MarketNotFound         = "Market not found"
	InternalServerError    = "Oops! Something went wrong..."
)

type Error struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

func apiError(error string, description string) Error {
	return Error{
		error,
		description,
	}
}
