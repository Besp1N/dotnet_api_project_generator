package dotnetproject

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
)

type DotnetProject struct {
	Name          string
	Dependencies  []string
	NuGetPackages []string
}

func New(name string) *DotnetProject {
	return &DotnetProject{
		Name: name,
	}
}

func runDotnetCommand(ctx context.Context, workingDir string, args ...string) error {
	cmd := exec.CommandContext(ctx, "dotnet", args...)
	cmd.Dir = workingDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("failed to run dotnet command: " + string(out))
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

func GenerateDotnetSolution(ctx context.Context, solutionName string, location string) error {
	err := runDotnetCommand(ctx, location, "new", "sln", "-n", solutionName)
	if err != nil {
		return err
	}

	return nil
}
