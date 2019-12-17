package main

import (
	"strings"
	"io/ioutil"
)

// HTMLWriter writes to an HTML file
type HTMLWriter struct {
	fileName string
}

// NewHTMLWriter creates a new value of type HTMLWriter
func NewHTMLWriter(fileName string) *HTMLWriter {
	if !strings.HasSuffix(fileName, ".html") {
		fileName = fileName+".html"
	}
	return &HTMLWriter{
		fileName: fileName,
	}
}

// Write writes the bytes to an HTML file
func (writer *HTMLWriter) Write(bytes []byte) (int, error)  {
	err := ioutil.WriteFile(writer.fileName, bytes, 0666)
	if err != nil {
		return 0, err
	}

	return  len(bytes), nil
}
