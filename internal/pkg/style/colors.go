package style

import (
	"encoding/json"
	"time"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
)

// ColorJSON formats the json with colors for the terminal
func ColorJSON(content string) string {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(content), &obj)
	if err != nil {
		eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed to display as JSON: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
		return content
	}
	jsonFormatter := colorjson.NewFormatter()
	jsonFormatter.Indent = 2
	s, err := jsonFormatter.Marshal(obj)
	if err != nil {
		return content
	}
	return string(s)
}

// Subtle use magenta and faint to format the text
func Subtle(s string) string {
	return color.New(color.FgMagenta, color.Faint).Sprint(s)
}

// Separator use magenta and faint to format the text
func Separator(s string) string {
	return color.New(color.FgBlack, color.Faint, color.Concealed).Sprint(s)
}

// Title make the text bold
func Title(s string) string {
	return color.New(color.Bold).Sprint(s)
}

// Loading make the text red and blink
func Loading(s string) string {
	return color.New(color.BlinkSlow, color.FgRed).Sprint(s)
}

// Completed make the text green
func Completed(s string) string {
	return color.New(color.FgGreen).Sprint(s)
}

// Header make the background blue and text white
func Header(s string) string {
	return color.New(color.BgBlue, color.FgWhite).Sprint(s)
}
