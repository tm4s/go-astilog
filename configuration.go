package astilog

import "flag"

// Flags
var (
	AppName  = flag.String("logger-app-name", "", "the logger's app name")
	Filename = flag.String("logger-filename", "", "the logger's filename")
	Verbose bool
)

// Formats
const (
	FormatJSON = "json"
	FormatText = "text"
)

// Outs
const (
	OutFile   = "file"
	OutStdOut = "stdout"
	OutSyslog = "syslog"
)

// Configuration represents the configuration of the logger
type Configuration struct {
	AppName         string `toml:"app_name"`
	DisableColors   bool   `toml:"disable_colors"`
	Filename        string `toml:"filename"`
	FullTimestamp   bool   `toml:"full_timestamp"`
	Format          string `toml:"format"`
	MessageKey      string `toml:"message_key"`
	Out             string `toml:"out"`
	TimestampFormat string `toml:"timestamp_format"`
	Verbose         bool   `toml:"verbose"`
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	if flag.Lookup("v") == nil {
		flag.BoolVar(&Verbose, "v", false, "if true, then log level is debug")
	} else {
		Verbose = flag.Lookup("v").Value.(flag.Getter).Get().(bool)
	}
	return Configuration{
		AppName:  *AppName,
		Filename: *Filename,
		Verbose:  Verbose,
	}
}
