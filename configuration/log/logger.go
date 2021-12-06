package log

type Logger interface {
	Info(message string)
	Error(message string, err error)
}
