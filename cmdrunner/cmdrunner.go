package cmdrunner

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Run(ctx context.Context, dir string, name string, args ...string) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(
			"Command failed: %s %s\nDir: %s\nError: %v\nOutput:\n%s",
			name, strings.Join(args, " "), dir, err, string(out),
		)
	}

	fmt.Printf("OK: %s %s\n", name, strings.Join(args, " "))
}
