package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Cmd struct {
	rootCmd *cobra.Command
}

func NewCmd() *Cmd {
	rootCmd := &cobra.Command{
		Use: "app",
	}

	rootCmd.AddCommand(
		NewGenerateCommand(),
		NewReadCommand(),
		NewRead2Command(),
		NewConvertCommand(),
		NewTestCommand(),
		NewConvert2Command(),
		NewConvert3Command(),
	)
	return &Cmd{
		rootCmd: rootCmd,
	}
}

func (cmd *Cmd) Execute() {
	if err := cmd.rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}

	os.Exit(0)
}
