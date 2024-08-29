package mylogger;

type Loggerr interface{
	Debug(format string,a ...interface{})
	Warning(format string,a ...interface{})
	Err(format string,a ...interface{})
}