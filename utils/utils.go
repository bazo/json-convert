package utils

import (
	"fmt"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func GetMemUsage() string {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// return fmt.Sprintf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
	// 	bToMb(memStats.Alloc), bToMb(memStats.TotalAlloc), bToMb(memStats.Sys), memStats.NumGC)

	return fmt.Sprintf("Alloc = %v MiB,  Sys = %v MiB, NumGC = %v\n",
		bToMb(memStats.Alloc), bToMb(memStats.Sys), memStats.NumGC)
}

func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func DockerPath(file string) string {
	return "files/" + file
}
