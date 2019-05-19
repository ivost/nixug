package services

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// RunWithEnv ...
func RunWithEnv(environ *[]string, command string, verbose bool, args ...string) (string, error) {
	if verbose {
		fmt.Printf("%s ", command)
		for _, v := range args {
			fmt.Printf("%s ", v)
		}
		println("")
	}
	cmd := exec.Command(command, args...)
	if environ != nil && len(*environ) > 0 {
		env := os.Environ()
		for _, pair := range *environ {
			env = append(env, pair)
		}
		cmd.Env = env
	}

	out, err := cmd.CombinedOutput()
	if err != nil && verbose {
		fmt.Printf("error %s ", err)
	}
	l := len(out)
	if l > 1 && out[l-1] == '\n' {
		// remove last element
		out = out[:l-1]
	}
	if verbose {
		fmt.Println(string(out))
	}
	return string(out), err
}

// Run runs command with variable number of args,
// returns output and error
func Run(command string, verbose bool, args ...string) (string, error) {
	return RunWithEnv(nil, command, verbose, args...)
}

func DelaySec(seconds int) {
	for i := 0; i < seconds; i++ {
		time.Sleep(1 * time.Second)
	}
}

//func MinimumDelaySec(minSec int, seconds int) {
//	if seconds < minSec {
//		seconds = minSec
//	}
//	DelaySec(seconds)
//}
