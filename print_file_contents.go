package main

import (
	"os"
	"fmt"
	"io"
)

// FilePrinter x
type FilePrinter struct {
	file *os.File
}

// NewFilePrinterFromArgs x
func NewFilePrinterFromArgs() FilePrinter {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	}

	return FilePrinter{
		file: file,
	}
}

// Print x
func (filePrinter FilePrinter) Print() {
	_, err := io.Copy(os.Stdout, filePrinter.file)
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(2)
	}
}
