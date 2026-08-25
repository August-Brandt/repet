package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/August-Brandt/repet/config"
	"github.com/spf13/cobra"
	"gopkg.in/alessio/shellescape.v1"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run the selected commands",
	Long:  `Run the selected commands directly`,
	RunE:  execute,
}

var repeat *bool

func _execute(in io.ReadCloser, out io.Writer) (err error) {
	flag := config.Flag

	var command string
	if (*repeat) {
		configDir, err := config.GetDefaultConfigDir()
		if err != nil {
			return err
		}
		lastCommandFilePath := path.Join(configDir, "last_command")
		if _, err := (os.Stat(lastCommandFilePath)); err != nil {
			fmt.Fprintln(out, "No previous command was stored on the system")
			return nil
		}
		commandBytes, err := os.ReadFile(lastCommandFilePath)
		if err != nil {
			return err
		}
		command = string(commandBytes)
	} else {
		var options []string
		if flag.Query != "" {
			options = append(options, fmt.Sprintf("--query %s", shellescape.Quote(flag.Query)))
		}
	
		commands, err := filter(options, flag.FilterTag, false)
		if err != nil {
			return err
		}
		command = strings.Join(commands, "; ")
	}
	// Show final command before executing it
	if !flag.Silent {
		fmt.Fprintf(out, "> %s\n", command)
	}

	return run(command, in, out)
}

func execute(cmd *cobra.Command, args []string) error {
	return _execute(os.Stdin, os.Stdout)
}

func init() {
	RootCmd.AddCommand(execCmd)
	execCmd.Flags().BoolVarP(&config.Flag.Color, "color", "", false,
		`Enable colorized output (only fzf)`)
	execCmd.Flags().StringVarP(&config.Flag.Query, "query", "q", "",
		`Initial value for query`)
	execCmd.Flags().StringVarP(&config.Flag.FilterTag, "tag", "t", "",
		`Filter tag`)
	execCmd.Flags().BoolVarP(&config.Flag.Silent, "silent", "s", false,
		`Suppress the command output`)
	execCmd.Flags().BoolVarP(&config.Flag.Copy, "copy", "c", false, 
		`Copies executed command to clipboard`)
	repeat = execCmd.Flags().BoolP("repeat", "r", false,
		`Repeats the previously executed command
This flag is mutually exclusive with query, color, and tag`)
}
