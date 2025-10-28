package golang

import "time"

// NowString returns current time formatted as yyyy-MM-dd HH:mm:ss
func NowString() string {
    return time.Now().Format("2006-01-02 15:04:05")
}