package exception

type BadRequest struct {
	Error string
}

func NewBadRequest(error string) BadRequest {
	return BadRequest{Error: error}
}
