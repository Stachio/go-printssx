package printssx

import "fmt"

// Noise - Enum
type Noise int

// Quiet - No output from the interface
const Quiet Noise = 0

//Subtle - Minimal amount of output from the interface
const Subtle Noise = 1

//Moderate - Subjective amount of output from the interface
const Moderate Noise = 2

//Loud - An obnoxious amount of output from the interface
const Loud Noise = 3

type printfx func(format string, a ...interface{})
type printlnx func(a ...interface{})

// Printer - Printer object for multi-use printing
type Printer struct {
	header       string
	verboseLevel Noise
	logLevel     Noise
	printfssx    printfx
	printlnssx   printlnx
}

// GetHeader - Self explanatory
func (printer *Printer) GetHeader() string {
	return printer.header
}

// GetVerboseLevel - Self explanatory
func (printer *Printer) GetVerboseLevel() Noise {
	return printer.verboseLevel
}

func (printer *Printer) canSpeak(level Noise) bool {
	return level <= printer.verboseLevel
}

func (printer *Printer) getHeaderStr(tag string) string {
	return fmt.Sprintf("[%s-%s]", printer.header, tag)
}

//Printf - Exported printf functionality
func (printer *Printer) Printf(level Noise, format string, a ...interface{}) {
	if printer.canSpeak(level) {
		format = printer.getHeaderStr("F") + " " + format
		printer.printfssx(fmt.Sprintf(format, a...))
	}
}

//Println - Exported Println functionality
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

//Errorf - Exported Errorf functionality
func (printer *Printer) Errorf(format string, a ...interface{}) error {
	format = printer.getHeaderStr("E") + " " + format
	return fmt.Errorf(format, a...)
}

//SetVerboseLevel - Set the verbosity level
func (printer *Printer) SetVerboseLevel(level Noise) {
	printer.verboseLevel = level
}

//New - Create a new Printer object
func New(header string, printlnssx printlnx, printfssx printfx, verboseLevel, logLevel Noise) *Printer {
	return &Printer{
		header:       header,
		printlnssx:   printlnssx,
		printfssx:    printfssx,
		verboseLevel: verboseLevel,
		logLevel:     logLevel,
	}
}
