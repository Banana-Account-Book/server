package httpCode

type Status struct {
	Code    int
	Message string
}

var (
	Ok                  = Status{Code: 200, Message: "Ok"}
	Created             = Status{Code: 201, Message: "Created"}
	Accepted            = Status{Code: 202, Message: "Accepted"}
	NoContent           = Status{Code: 204, Message: "No Content"}
	BadRequest          = Status{Code: 400, Message: "Bad Request"}
	Unauthorized        = Status{Code: 401, Message: "Unauthorized"}
	Forbidden           = Status{Code: 403, Message: "Forbidden"}
	NotFound            = Status{Code: 404, Message: "Not Found"}
	Conflict            = Status{Code: 409, Message: "Conflict"}
	InternalServerError = Status{Code: 500, Message: "Internal Server Error"}
	NotImplemented      = Status{Code: 501, Message: "Not Implemented"}
)
