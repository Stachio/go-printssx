package printssx

import (
	"fmt"
	"strings"
)

// Noise - Enum
type Noise int

// Quiet - No output from the interface
const Quiet Noise = 0

// Subtle - Minimal amount of output from the interface
const Subtle Noise = 1

// Moderate - Subjective amount of output from the interface
const Moderate Noise = 2

// Loud - An obnoxious amount of output from the interface
const Loud Noise = 3

type printfx func(format string, a ...interface{})
type printlnx func(a ...interface{})

// Printer - Printer object for multi-use printing
type Printer struct {
	headers      []string
	verboseLevel Noise
	logLevel     Noise
	printfssx    printfx
	printlnssx   printlnx
}

// GetHeaders - Returns the array of headers
func (printer *Printer) GetHeaders() []string {
	headers := make([]string, len(printer.headers))
	for i, header := range printer.headers {
		headers[i] = header
	}
	return headers
}

// GetHeaders - Independent GetHeaders functionality
func GetHeaders(printer *Printer) []string {
	return printer.GetHeaders()
}

// GetVerboseLevel - Self explanatory
func (printer *Printer) GetVerboseLevel() Noise {
	return printer.verboseLevel
}

func (printer *Printer) canSpeak(level Noise) bool {
	return level <= printer.verboseLevel
}

func (printer *Printer) getHeaderStr(tag string) string {
	headerStrs := make([]string, len(printer.headers))
	for i, header := range printer.headers {
		headerStrs[i] = "[" + header + "]"
	}
	return strings.Join(headerStrs, "") + "[" + tag + "]"
}

// Printf - Exported printf functionality
func (printer *Printer) Printf(level Noise, format string, a ...interface{}) {
	if printer.canSpeak(level) {
		format = printer.getHeaderStr("F") + " " + format
		printer.printfssx(fmt.Sprintf(format, a...))
	}
}

// Printf - Independent Printf functionality
func Printf(printer *Printer, level Noise, format string, a ...interface{}) {
	printer.Printf(level, format, a...)
}

// Println - Exported Println functionality
func (printer *Printer) Println(level Noise, a ...interface{}) {
	if printer.canSpeak(level) {
		aa := make([]interface{}, len(a)+1)
		aa[0] = printer.getHeaderStr("L")
		for i := range a {
			aa[i+1] = a[i]
		}
		printer.printlnssx(aa...)
	}
}

// Println - Independent Println functionality
func Println(printer *Printer, level Noise, a ...interface{}) {
	printer.Println(level, a...)
}

// Errorf - Exported Errorf functionality
func (printer *Printer) Errorf(format string, a ...interface{}) error {
	format = printer.getHeaderStr("E") + " " + format
	return fmt.Errorf(format, a...)
}

// Errorf - Independent Errorf functionality
func Errorf(printer *Printer, format string, a ...interface{}) error {
	return printer.Errorf(format, a...)
}

// SetVerboseLevel - Sets the verbosity level
func (printer *Printer) SetVerboseLevel(level Noise) {
	printer.verboseLevel = level
}

// SetVerboseLevel - Independent SetVerboseLevel functionality
func SetVerboseLevel(printer *Printer, level Noise) {
	printer.SetVerboseLevel(level)
}

// PushHeader - Pushes a new header onto the header stack
func (printer *Printer) PushHeader(header string) {
	printer.headers = append(printer.headers, header)
}

// PopHeader - Pops the last header off the header stack
// For safety purposes, the input header must match the last header on the stack
func (printer *Printer) PopHeader(header string) {
	lastHeader := printer.headers[len(printer.headers)-1]
	if header != lastHeader {
		panic(fmt.Errorf("Last header \"%s\" does not match input header \"%s\"", lastHeader, header))
	} else if len(printer.headers) == 1 {
		panic(fmt.Errorf("Unable to pop first header \"%s\" from stack", lastHeader))
	}
	printer.headers = printer.headers[:len(printer.headers)-1]
}

// New - Creates a new Printer object
func New(header string, printlnssx printlnx, printfssx printfx, verboseLevel, logLevel Noise) *Printer {
	return &Printer{
		headers:      []string{header},
		printlnssx:   printlnssx,
		printfssx:    printfssx,
		verboseLevel: verboseLevel,
		logLevel:     logLevel,
	}
}
