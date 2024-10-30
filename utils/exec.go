package utils

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

func ExecCommand(command string, arg ...string) error {
	start := time.Now().UnixMilli()
	cmd := exec.Command(command, arg...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	cost := time.Now().UnixMilli() - start
	if err != nil {
		if cmd.ProcessState.ExitCode() == 200 {
			log.Printf("execute command %s ok, cost %dms", command, cost)
			return nil
		}
		log.Printf("execute command %s error: %v\nstdout: %s\nstderr: %s", command, err, out.String(), stderr.String())
		return err
	}

	return nil
}

func ExecCommandWithTimeout(command string, timeout time.Duration, arg ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	start := time.Now().UnixMilli()
	cmd := exec.CommandContext(ctx, command, arg...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	cost := time.Now().UnixMilli() - start
	if err != nil {
		if cmd.ProcessState.ExitCode() == 200 {
			log.Printf("execute command %s ok, cost %dms", command, cost)
			return nil
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("execute command %s timeout, cost %dms", command, cost)
			return context.DeadlineExceeded
		}
		log.Printf("execute command %s error: %v\nstdout: %s\nstderr: %s", command, err, out.String(), stderr.String())
		return err
	}

	return nil
}
