package main

import (
	"context"
	"dotnetApiGenerator/cmdrunner"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	ctx := context.Background()

	projectName := flag.String("name", "", "Name of the solution (base name for projects)")
	location := flag.String("location", "", "Output location for the solution")
	flag.Parse()

	if *projectName == "" {
		log.Fatal("missing required flag: --name")
	}
	if *location == "" {
		log.Fatal("missing required flag: --location")
	}

	slnDir := filepath.Join(*location, *projectName)
	if err := os.MkdirAll(slnDir, 0o755); err != nil {
		log.Fatalf("Cannot create directory %s: %v", slnDir, err)
	}

	slnFile := *projectName + ".sln"

	domainName := *projectName + ".Domain"
	appName := *projectName + ".Application"
	infraName := *projectName + ".Infrastructure"
	apiName := *projectName + ".Api"
	testsName := *projectName + ".Tests"

	domainCsproj := filepath.Join(domainName, domainName+".csproj")
	appCsproj := filepath.Join(appName, appName+".csproj")
	infraCsproj := filepath.Join(infraName, infraName+".csproj")
	apiCsproj := filepath.Join(apiName, apiName+".csproj")
	testsCsproj := filepath.Join(testsName, testsName+".csproj")

	fmt.Println("Starting generation process...")

	fmt.Println("Creating solution...")
	cmdrunner.Run(ctx, slnDir, "dotnet", "new", "sln", "-n", *projectName)
	fmt.Println("Solution created.")

	fmt.Println("Creating projects...")
	cmdrunner.Run(ctx, slnDir, "dotnet", "new", "classlib", "-n", domainName, "-o", domainName)
	cmdrunner.Run(ctx, slnDir, "dotnet", "new", "classlib", "-n", appName, "-o", appName)
	cmdrunner.Run(ctx, slnDir, "dotnet", "new", "classlib", "-n", infraName, "-o", infraName)
	cmdrunner.Run(ctx, slnDir, "dotnet", "new", "webapi", "-n", apiName, "-o", apiName)
	cmdrunner.Run(ctx, slnDir, "dotnet", "new", "xunit", "-n", testsName, "-o", testsName)
	fmt.Println("Projects created.")

	fmt.Println("Setting up solution and project references...")
	cmdrunner.Run(ctx, slnDir, "dotnet", "sln", slnFile, "add", domainCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "sln", slnFile, "add", appCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "sln", slnFile, "add", infraCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "sln", slnFile, "add", apiCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "sln", slnFile, "add", testsCsproj)

	cmdrunner.Run(ctx, slnDir, "dotnet", "add", appCsproj, "reference", domainCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "add", infraCsproj, "reference", appCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "add", infraCsproj, "reference", domainCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "add", apiCsproj, "reference", appCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "add", apiCsproj, "reference", infraCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "add", testsCsproj, "reference", appCsproj)
	cmdrunner.Run(ctx, slnDir, "dotnet", "add", testsCsproj, "reference", domainCsproj)
	fmt.Println("Solution and project references set up.")

	fmt.Println("Done.")
}
