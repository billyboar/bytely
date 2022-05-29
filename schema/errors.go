package schema

const (
	ErrBadPayload       = "BAD_PAYLOAD"
	ErrGRPCServiceError = "GRPC_SERVICE_ERROR"
	ErrJSONMarshalError = "JSON_MARSHAL_ERROR"
	ErrNotFound         = "RESOURCE_NOT_FOUND"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
