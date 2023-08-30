package video_stream

import "fmt"

func StreamFormat(formatName string) string {
	format := fmt.Sprintf("Processing %v format", formatName)
	return format
}
