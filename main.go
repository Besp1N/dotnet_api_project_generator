package main

import (
	"dotnetApiGenerator/dotnetproject"
	"fmt"
)

func main() {
	fmt.Print("Enter project name: ")

	var projectName string
	_, err := fmt.Scanln(&projectName)
	if err != nil {
		panic("failed to read project name input")
	}

	var location string
	fmt.Print("Enter location: ")
	_, err = fmt.Scanln(&location)
	if err != nil {
		panic("failed to read location input")
	}

	fmt.Println("Starting generation process...")
	fmt.Println()

	projectNames := dotnetproject.GenerateProjectNames(projectName)
	dotnetProjects := make([]*dotnetproject.DotnetProject, 0, len(projectNames))

	fmt.Println(len(projectNames), "projects will be created in", location, ":")
	for _, name := range projectNames {
		fmt.Println("- " + name)
		dotnetProject := dotnetproject.New(name, location)
		dotnetProjects = append(dotnetProjects, dotnetProject)
	}

	fmt.Println()

}
