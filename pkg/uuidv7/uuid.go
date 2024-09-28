package uuidv7

import (
	"github.com/google/uuid"
)

// GenerateUUIDv7 generates a UUID version 7.
// It returns the generated UUID as a uuid.UUID type and any error encountered.
// Example usage:
//
//	id, err := GenerateUUIDv7()
//	if err != nil {
//	    log.Fatal("Failed to generate UUID:", err)
//	}
//	fmt.Println("Generated UUID v7:", id)
func GenerateUUIDv7() (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// GenerateUUIDv7String generates a UUID version 7 and returns it as a string.
// It calls GenerateUUIDv7 to get the UUID and converts it to a string format.
// Example usage:
//
//	uuidStr, err := GenerateUUIDv7String()
//	if err != nil {
//	    log.Fatal("Failed to generate UUID string:", err)
//	}
//	fmt.Println("Generated UUID v7 as string:", uuidStr)
func GenerateUUIDv7String() (string, error) {
	id, err := GenerateUUIDv7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// ToUUIDv7 converts a string representation of a UUID to a uuid.UUID.
// It returns the parsed UUID and an error if the string is not a valid UUID.
// Example usage:
//
//	uuidStr := "your-uuid-string-here"
//	id, err := ToUUIDv7(uuidStr)
//	if err != nil {
//	    log.Fatal("Invalid UUID string:", err)
//	}
//	fmt.Println("Parsed UUID v7:", id)
func ToUUIDv7(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// MustParseToUUIDv7 converts a string to a uuid.UUID and panics if the string is invalid.
// This function is useful when you are sure the string is a valid UUID and want to avoid error handling.
// Example usage:
//
//	uuidStr := "your-uuid-string-here"
//	id := MustParseToUUIDv7(uuidStr) // This will panic if invalid
//	fmt.Println("Must parsed UUID v7:", id)
func MustParseToUUIDv7(s string) uuid.UUID {
	u, _ := uuid.Parse(s) // Ignoring error, as this function will panic on invalid input
	return u
}

// UUIDv7ToString converts a uuid.UUID to its string representation.
// This is useful when you need to store or display the UUID in string format.
// Example usage:
//
//	id := uuid.New() // Replace with actual UUID generation
//	uuidStr := UUIDv7ToString(id)
//	fmt.Println("UUID v7 as string:", uuidStr)
func UUIDv7ToString(id uuid.UUID) string {
	return id.String()
}
