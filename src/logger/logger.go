package logger

import (
	"fmt"
	"time"
)

const (
	Log_level_Normal = 3
	Log_level_Debug  = 2
	Log_level_Error  = 1
)

var Log_level int = 3

func SetLogLevel(level int) {
	Log_level = level
}
func LogErrorF(err error, info string) {
	if err != nil && Log_level == Log_level_Error {
		t := time.Now()
		year := t.Year()
		month := t.Month().String()
		day := t.Day()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		fmt.Printf("%d-%s-%d %d:%d:%d ERROR : %s , %s\n", year, month, day, h, m, s, err.Error(), info)
	}
}
func LogError(info string) {
	if Log_level == Log_level_Error {
		t := time.Now()
		year := t.Year()
		month := t.Month().String()
		day := t.Day()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		fmt.Printf("%d-%s-%d %d:%d:%d  ERROR : %s \n", year, month, day, h, m, s, info)
	}
}
func LogNormalF(err error, info string) {
	if err != nil && Log_level >= Log_level_Normal {
		t := time.Now()
		year := t.Year()
		month := t.Month().String()
		day := t.Day()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		fmt.Printf("%d-%s-%d %d:%d:%d NORMAL: %s , %s\n", year, month, day, h, m, s, err.Error(), info)
	}
}
func LogNormal(info string) {
	if Log_level >= Log_level_Normal {
		t := time.Now()
		year := t.Year()
		month := t.Month().String()
		day := t.Day()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		fmt.Printf("%d-%s-%d %d:%d:%d NORMAL: %s\n", year, month, day, h, m, s, info)
	}
}
func LogDebugF(err error, info string) {
	if err != nil && Log_level == Log_level_Debug {
		t := time.Now()
		year := t.Year()
		month := t.Month().String()
		day := t.Day()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		fmt.Printf("%d-%s-%d %d:%d:%d DEBUG : %s , %s\n", year, month, day, h, m, s, err.Error(), info)
	}
}
func LogDebug(info string) {
	if Log_level == Log_level_Debug {
		t := time.Now()
		year := t.Year()
		month := t.Month().String()
		day := t.Day()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		fmt.Printf("%d-%s-%d %d:%d:%d DEBUG : %s\n", year, month, day, h, m, s, info)
	}
}
func Log(info string) {
	if Log_level == Log_level_Error {
		LogError(info)
	}
	if Log_level == Log_level_Debug {
		LogDebug(info)

	}
	if Log_level >= Log_level_Normal {
		LogNormal(info)
	}
}
func LogF(err error, info string) {
	if Log_level == Log_level_Error {
		LogErrorF(err, info)
	}
	if Log_level == Log_level_Debug {
		LogDebugF(err, info)

	}
	if Log_level >= Log_level_Normal {
		LogNormalF(err, info)
	}
}
