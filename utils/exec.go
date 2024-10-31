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
			log.Printf("Execute command %s with args %v ok, cost %dms", command, args, cost)
			return string(output), nil
		}

		log.Printf("Execute command %s with args %v error: %v\noutput: %s", command, args, err, string(output))
		return "", err
	}

	log.Printf("Execute command %s with args %v success, cost %dms", command, args, cost)
	return string(output), nil
}

func ExecCommandContext(ctx context.Context, command string, args ...string) (string, error) {
	start := time.Now()
	cmd := exec.CommandContext(ctx, command, args...)

	output, err := cmd.CombinedOutput()
	cost := time.Since(start).Milliseconds()

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Execute command %s with args %v timeout, cost %dms", command, args, cost)
			return string(output), context.DeadlineExceeded
		}

		exitCode := -1
		if cmd.ProcessState != nil {
			exitCode = cmd.ProcessState.ExitCode()
		}

		if exitCode == 200 {
			log.Printf("Execute command %s with args %v ok, cost %dms", command, args, cost)
			return string(output), nil
		}

		log.Printf("Execute command %s with args %v error: %v\noutput: %s", command, args, err, string(output))
		return "", err
	}

	log.Printf("Execute command %s with args %v success, cost %dms", command, args, cost)
	return string(output), nil
}
