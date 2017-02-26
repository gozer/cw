package main

import (
	"time"

	"github.com/lucagrulla/cloudwatch-tail/cloudwatch"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	timeFormat = "2006-01-02T15:04:05"

	tailCommand     = kingpin.Command("tail", "Tail a log group")
	lsCommand       = kingpin.Command("ls", "show all log groups")
	logGroupPattern = lsCommand.Arg("group", "the log group name").String()
	follow          = tailCommand.Flag("follow", "don't stop when the end of stream is reached").Short('f').Default("false").Bool()
	logGroupName    = tailCommand.Arg("group", "The log group name").Required().String()
	startTime       = tailCommand.Arg("start", "The start time").Default(time.Now().Format(timeFormat)).String()
	streamName      = tailCommand.Arg("stream", "Stream name").String()
)

func formatTimestamp(ts int64) string {
	return time.Unix(ts, 0).Format(timeFormat)
}

func main() {
	kingpin.Version("0.0.1")
	command := kingpin.Parse()

	switch command {
	case "ls":
		cloudwatch.Ls()
	case "tail":
		cloudwatch.Tail(startTime, follow, logGroupName, streamName)
	}
}
