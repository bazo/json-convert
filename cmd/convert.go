package cmd

import (
	"bufio"
	"encoding/json"
	"json-convert/types2"
	"json-convert/utils"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	"github.com/parquet-go/parquet-go"
)

func NewConvertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "convert [file]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			infile := filename + ".jsonl"

			outfile := filename + ".parquet"

			f, err := os.Open(utils.DockerPath(infile))
			if err != nil {
				log.Fatalln("Cannot open file", infile)
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)

			stats, err := f.Stat()

			if err != nil {
				log.Fatalln("Cannot stat file", infile)
			}

			fw, err := os.Create(utils.DockerPath(outfile))
			if err != nil {
				log.Println("Can't create local file", err)
				return
			}
			defer fw.Close()

			pw := parquet.NewWriter(fw)

			bar := progressbar.DefaultBytes(stats.Size())
			for scanner.Scan() {
				b := scanner.Bytes()
				line := &types2.Line{}
				json.Unmarshal(b, line)

				err = pw.Write(line.ToParquet())
				if err != nil {
					log.Fatalln(err)
				}

				bar.Add(len(b))
				bar.Describe(utils.GetMemUsage())
			}

			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

			err = pw.Close()
			if err != nil {
				log.Fatal("Error stopping writer:", err)
			}

			bar.Finish()
		},
	}

	return cmd
}
