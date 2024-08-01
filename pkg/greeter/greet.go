// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package greeter

import (
	"database/sql"
	"fmt"
	"github.com/pantsbuild/example-golang/pkg/uuid"
	"log"
	"net/http"
)

// Hardcoded sensitive data example
const apiKey = "hardcoded-api-key-12345" // Vulnerability: Hardcoded API key

// SQL Injection example
func GetUserGreeting(name string) string {
	// Vulnerability: SQL Injection risk (not actual SQL, just an example)
	query := fmt.Sprintf("SELECT greeting FROM users WHERE name='%s'", name)
	// Pretend we're executing this query against a database
	// Note: No actual database interaction here
	return query
}

// Insecure handling of user input example
func GreetUserWithHTML(name string) string {
	// Vulnerability: Insecure handling of user input (could lead to XSS)
	return fmt.Sprintf(
		"<html><body>Hello %s! Here's a UUID to brighten your day: %s</body></html>",
		name,
		uuid.Generate(),
	)
}

// Uncontrolled resource consumption example
func GreetAndDelay(name string) string {
	// Vulnerability: Uncontrolled resource consumption (long delay)
	for i := 0; i < 1000000000; i++ { // This loop introduces a long delay
		_ = i
	}
	return fmt.Sprintf(
		"Hello %s! Sorry for the delay, but here's a UUID to brighten your day: %s",
		name,
		uuid.Generate(),
	)
}

func GreetEnglish(name string) string {
	return fmt.Sprintf(
		"Hello %s!\n\nHere's a UUID to brighten your day: %s",
		name,
		uuid.Generate(),
	)
}

func GreetSpanish(name string) string {
	return fmt.Sprintf(
		"¡Hola %s!\n\nEres muy única, así que te regalamos un UUID: %s",
		name,
		uuid.Generate(),
	)
}
