package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"json-convert/types"
	"log"
	"os"
	"runtime"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func getMemUsage() string {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// return fmt.Sprintf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
	// 	bToMb(memStats.Alloc), bToMb(memStats.TotalAlloc), bToMb(memStats.Sys), memStats.NumGC)

	return fmt.Sprintf("Alloc = %v MiB,  Sys = %v MiB, NumGC = %v\n",
		bToMb(memStats.Alloc), bToMb(memStats.Sys), memStats.NumGC)
}

func NewReadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "read [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			file := filename + ".jsonl"

			f, err := os.Open(file)
			if err != nil {
				log.Fatalln("Cannot open file", file)
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)

			stats, err := f.Stat()

			if err != nil {
				log.Fatalln("Cannot stat file", file)
			}

			bar := progressbar.DefaultBytes(stats.Size())
			for scanner.Scan() {
				b := scanner.Bytes()
				line := &types.Line{}
				json.Unmarshal(b, line)

				//log.Println(line)

				bar.Add(len(b))
				bar.Describe(getMemUsage())
				b = nil
				line = nil
			}

			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}
			bar.Finish()
		},
	}

	return cmd
}
