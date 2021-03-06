package logging

import (
	"fmt"
	"runtime"
	"time"
	"strconv"
)

type NewLogger struct {
	debugLevel int
	loggerOpen bool
	appName string
}

//----------- Public Functions ----------------------

func (l *NewLogger) Open() {

	l.loggerOpen = true
}

func (l *NewLogger) Close() {

	l.loggerOpen = false
}
 
func (l *NewLogger) Debug(message string) {

	l.processMessage(message, "Debug")
}

func (l *NewLogger) Info(message string) {
	
	l.processMessage(message, "Info")
}

func (l *NewLogger) Warning(message string) {

	l.processMessage(message, "Warning")
}

func (l NewLogger) Error(message string) {

	l.processMessage(message, "Error")
}
 

func (l *NewLogger) SetDebugLevel(debugLevel int) {

	l.debugLevel = debugLevel
}

func (l *NewLogger) GetDebugLevel() int{

	return l.debugLevel
}

func (l *NewLogger) SetAppName(name string) {

	l.appName = name
}

func (l *NewLogger) GetAppName() string{

	return l.appName
}

func (l *NewLogger) IsOpen() bool{

	return l.loggerOpen
}

//-------------Private Functions ------------------------------

func (l *NewLogger) processMessage(message string, logLevelName string){

	if l.loggerOpen == false {
		 panic("A logger function was called and the logger is currently closed, call 'log.Open' before calling logger functions.")
	}

	
	//Get the caller source name and line number
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	fileFunc, lineFunc := f.FileLine(pc)

	//get the local time stamp
	t := time.Now().Local()
	//var timeStamp = t.Format("20060102150405")

	var messageToSave = t.String() + "|" + logLevelName + "|"+ l.appName + "|" + message + "|" + fileFunc+ ":" + strconv.Itoa(lineFunc)

	fmt.Println(messageToSave)

}



