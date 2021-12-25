//Package log implements logs interfaces
package logger

//Logger is an interface that will log all the data that application will need to print
type Logger interface {
	Info(message string)
	Error(message string, err error)
}
