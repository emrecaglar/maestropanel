package maestropanel

type Result struct {
	StatusCode int32
	ErrorCode  int32
	Message    string
}

type OperationResult struct {
	Code    int32
	Message string
}
