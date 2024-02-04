package storm_go_logger

import (
	"log"
	"os"
	"regexp"
	"strings"
)

// Environment is an enumerated type representing the environment in which
// the application is running: Development or Production.
type Environment uint

const (
	Production Environment = iota
	Development
)

// GetEnv gets the environment in which the application is running from the "ENVIRONMENT" environment variable.
func GetEnv() Environment {
	environment := os.Getenv("ENVIRONMENT")
	// write a regular expression to match the production environment
	pattern, err := regexp.Compile("(production|prod|pro|release|rel|deploy)")
	if err != nil {
		// Use the regular log package to prevent circular imports
		log.Fatalln("could not compile regular expression")
	}
	var env Environment
	if pattern.MatchString(strings.ToLower(environment)) {
		env = Production
	} else {
		env = Development
	}
	return env
}
