package web

func apiError(msg string) struct {
	Message string `json:"message"`
} {
	return struct {
		Message string `json:"message"`
	}{
		msg,
	}
}
