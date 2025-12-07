package main

import (
	"context"
	"dotnetApiGenerator/dotnetproject"
	"flag"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	fmt.Print("Enter project name: ")

	projectName := flag.String("name", "", "Name of the project/solution")
	location := flag.String("location", "", "Output location for the solution")

	flag.Parse()

	if *projectName == "" {
		log.Fatal("missing required flag: --name")
	}

	if *location == "" {
		log.Fatal("missing required flag: --location")
	}

	fmt.Println("Starting generation process...")
	fmt.Println()

	projectNames := dotnetproject.GenerateProjectNames(*projectName)
	dotnetProjects := make([]*dotnetproject.DotnetProject, 0, len(projectNames))

	fmt.Println(len(projectNames), "projects will be created in", location, ":")
	for _, name := range projectNames {
		fmt.Println("- " + name)
		dotnetProject := dotnetproject.New(name)
		dotnetProjects = append(dotnetProjects, dotnetProject)
	}

	fmt.Println()

	err := dotnetproject.GenerateDotnetSolution(ctx, *projectName, *location)
	if err != nil {
		fmt.Println(err)
	}
}
