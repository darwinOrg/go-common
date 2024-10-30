package utils

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

func ExecCommand(command string, arg ...string) (string, error) {
	start := time.Now().UnixMilli()
	cmd := exec.Command(command, arg...)

	output, err := cmd.CombinedOutput()
	cost := time.Now().UnixMilli() - start
	if err != nil {
		if cmd.ProcessState.ExitCode() == 200 {
			log.Printf("execute command %s ok, cost %dms", command, cost)
			return string(output), nil
		}
		log.Printf("execute command %s error: %v\noutput: %s", command, err, string(output))
		return "", err
	}

	return string(output), nil
}

func ExecCommandWithTimeout(command string, timeout time.Duration, arg ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	start := time.Now().UnixMilli()
	cmd := exec.CommandContext(ctx, command, arg...)

	output, err := cmd.CombinedOutput()
	cost := time.Now().UnixMilli() - start
	if err != nil {
		if cmd.ProcessState.ExitCode() == 200 {
			log.Printf("execute command %s ok, cost %dms", command, cost)
			return string(output), nil
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("execute command %s timeout, cost %dms", command, cost)
			return string(output), context.DeadlineExceeded
		}
		log.Printf("execute command %s error: %v\noutput: %s", command, err, string(output))
		return "", err
	}

	return string(output), nil
}
