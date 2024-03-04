package internal

import "errors"

var (
	ERRNOTFOUND     = errors.New("Not Found")
	ERRINTERNAL     = errors.New("Internal Server Error")
	ERRBADREQUEST   = errors.New("Bad Request")
	ERRUNAUTHORIZED = errors.New("Unauthorized")
	ERRFORBIDDEN    = errors.New("Forbidden")
)
