package logs

import (
	"errors"
	"io"
	"log"
	"os"
	"runtime"
)

type Brush func(string) string

func NewBrush(color string) Brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []Brush{
	NewBrush("1;37"), // App    white (only for pholcus)
	NewBrush("1;37"), // Emergency  white
	NewBrush("1;36"), // Alert  cyan
	NewBrush("1;35"), // Critical   magenta
	NewBrush("1;31"), // Error  red
	NewBrush("1;33"), // Warning    yellow
	NewBrush("1;32"), // Notice green
	NewBrush("1;34"), // Informational  blue
	NewBrush("1;34"), // Debug  blue
}

// ConsoleWriter implements LoggerInterface and writes messages to terminal.
type ConsoleWriter struct {
	lg    *log.Logger
	Level int `json:"level"`
}

// create ConsoleWriter returning as LoggerInterface.
func NewConsole() LoggerInterface {
	cw := &ConsoleWriter{
		Level: LevelDebug,
		lg:    log.New(os.Stdout, "", log.LstdFlags),
	}
	return cw
}

// init console logger.
// config like map[string]interface{}{"level":LevelTrace,"writer":os.Stdout}.
func (c *ConsoleWriter) Init(config map[string]interface{}) error {
	if config == nil {
		return nil
	}
	if l, ok := config["level"]; ok {
		if l2, ok2 := l.(int); ok2 {
			c.Level = l2
		} else {
			return errors.New("consloe config-level's type is incorrect!")
		}
	}
	if w, ok := config["writer"]; ok {
		if w2, ok2 := w.(io.Writer); ok2 {
			c.lg = log.New(w2, "", log.LstdFlags)
		}
	}
	return nil
}

// write message in console.
func (c *ConsoleWriter) WriteMsg(msg string, level int) error {
	if level > c.Level {
		return nil
	}
	if goos := runtime.GOOS; goos == "windows" {
		c.lg.Println(msg)
		return nil
	}
	c.lg.Println(colors[level](msg))

	return nil
}

// implementing method. empty.
func (c *ConsoleWriter) Destroy() {

}

// implementing method. empty.
func (c *ConsoleWriter) Flush() {

}

func init() {
	Register("console", NewConsole)
}