//go:build !windows

package cmd

import (
	"io"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/August-Brandt/repet/config"
	repetPath "github.com/August-Brandt/repet/path"
	"github.com/atotto/clipboard"
)

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	if len(config.Conf.General.Cmd) > 0 {
		line := append(config.Conf.General.Cmd, command)
		cmd = exec.Command(line[0], line[1:]...)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	if config.Flag.Copy {
		clipboard.WriteAll(command)
	}
	if !config.Flag.Exclude {
		configDir, err := config.GetDefaultConfigDir()
		if err != nil {
			return err
		}
		err = os.WriteFile(path.Join(configDir, "last_command"), []byte(command), 0644)
		if err != nil {
			return err
		}
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

func editFile(command string, filePath repetPath.AbsolutePath, startingLine int) error {
	command += " +" + strconv.Itoa(startingLine) + " " + filePath.Get()
	return run(command, os.Stdin, os.Stdout)
}
