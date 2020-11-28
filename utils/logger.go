package utils

// Logger is the interface for logging. We use sirupsen/logrus by default, but,
// to make things agnostic (and because I want so), let's use a interface. Also,
// this interface forces structured logging, which is nice.
type Logger interface {
	// Debugf is the default debug message.
	Debugf(format string, args ...interface{})
	// Infof is the default debug message.
	Infof(format string, args ...interface{})
	// Warnf is the default debug message.
	Warnf(format string, args ...interface{})
	// Errorf is the default debug message.
	Errorf(format string, args ...interface{})
	// Fatalf is the default debug message.
	Fatalf(format string, args ...interface{})
	// Panicf is the default debug message.
	Panicf(format string, args ...interface{})
}
