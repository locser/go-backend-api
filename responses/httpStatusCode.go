package responses

const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
	ErrCodeParamInvalid       = 10001
)

// message
var Msg = map[int]string{
	StatusOK:                  "Success",
	StatusCreated:             "Created",
	StatusBadRequest:          "Bad Request",
	StatusUnauthorized:        "Unauthorized",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",

	ErrCodeParamInvalid: "Param Invalid",
}
