package dotnetproject

import (
	"context"
	"fmt"
	"os/exec"
)

type DotnetProject struct {
	Name          string
	Location      string
	Dependencies  []string
	NuGetPackages []string
}

func New(name string, location string) *DotnetProject {
	return &DotnetProject{
		Name:     name,
		Location: location,
	}
}

func runDotnetCommand(ctx context.Context, workingDir string, args ...string) error {
	cmd := exec.CommandContext(ctx, "dotnet", args...)
	cmd.Dir = workingDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic("Failed to run dotnet command: " + string(out))
	}

	return nil
}

func GenerateProjectNames(projectName string) []string {
	return []string{
		fmt.Sprintf("%s.Api", projectName),
		fmt.Sprintf("%s.Application", projectName),
		fmt.Sprintf("%s.Domain", projectName),
		fmt.Sprintf("%s.Infrastructure", projectName),
		fmt.Sprintf("%s.Worker", projectName),
	}
}
