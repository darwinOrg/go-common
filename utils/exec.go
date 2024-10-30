package utils

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

func ExecCommand(command string, args ...string) (string, error) {
	start := time.Now()
	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()
	cost := time.Since(start).Milliseconds()

	if err != nil {
		if cmd.ProcessState != nil && cmd.ProcessState.ExitCode() == 200 {
			log.Printf("Execute command %s ok, cost %dms", command, cost)
			return string(output), nil
		}

		log.Printf("Execute command %s error: %v\noutput: %s", command, err, string(output))
		return "", err
	}

	log.Printf("Execute command %s success, cost %dms", command, cost)
	return string(output), nil
}

func ExecCommandWithTimeout(command string, timeout time.Duration, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()
	cmd := exec.CommandContext(ctx, command, args...)

	output, err := cmd.CombinedOutput()
	cost := time.Since(start).Milliseconds()

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Execute command %s timeout, cost %dms", command, cost)
			return string(output), context.DeadlineExceeded
		}

		exitCode := -1
		if cmd.ProcessState != nil {
			exitCode = cmd.ProcessState.ExitCode()
		}

		if exitCode == 200 {
			log.Printf("Execute command %s ok, cost %dms", command, cost)
			return string(output), nil
		}

		log.Printf("Execute command %s error: %v\noutput: %s", command, err, string(output))
		return "", err
	}

	log.Printf("Execute command %s success, cost %dms", command, cost)
	return string(output), nil
}
