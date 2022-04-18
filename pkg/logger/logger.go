package logger

import (
	"io"
)

type Logger struct {
	outstream io.Writer
	channel <-chan []byte
}

func (l Logger) writeLogs() {
	for logs := range l.channel {
		l.outstream.Write(logs)
	}	
}

func (l Logger) Init(ostream io.Writer, channel <-chan []byte) {
	l.outstream = ostream
	l.channel = channel
	go l.writeLogs()
}