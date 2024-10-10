package queue

const (
	BasicError = "basic_error"
)

type BasicErrorMessage struct {
	Queue   string `json:"queue"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
