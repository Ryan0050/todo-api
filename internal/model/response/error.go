package response

var (
	ErrDataNotFound           = HttpError(404, 404, "Data Not Found")
	ErrBadRequest             = HttpError(400, 400, "Bad Request")
	ErrInternalServer         = HttpError(500, 500, "Internal Server Error")
	ErrDataAlreadyExists      = HttpError(400, 4001, "Data already exists")
	ErrUsernameAlreadyExists  = HttpError(400, 4002, "Username already in use")
	ErrUnauthorized           = HttpError(401, 401, "Unauthorized")
	ErrForbidden              = HttpError(403, 403, "Forbidden")
	ErrDuplicateID            = HttpError(400, 400, "Duplicate ID")
	ErrWrongIDRegexValidation = HttpError(400, 4003, "ID must contain only a-z (lowercase or uppercase) and '.'")
)
