package util

import (
	"egghead/app/common"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// GetEnv retrieves the value of the specified environment variable or returns a default value if not set
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	strVal := GetEnv(key, "")

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}

	return defaultValue
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseBool(strVal); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvWithEnum verifies is the allowed value is within the enums provided
func GetEnvWithEnum(key string, defaultValue string, allowedvalues []string) string {
	if !ContainsString(allowedvalues, defaultValue) {
		log.Panic("Default value is not in the allowed values list.")
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	if !ContainsString(allowedvalues, val) {
		log.Println("Value is not allowed. Fallback to default value.")
		return defaultValue
	}

	return val
}

// ConvertStrToInt converts a string to an integer, returns math.MaxInt if conversion fails
func ConvertStrToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
		return math.MaxInt
	}
	return i
}

// GenerateSlug generates a slug from the given string
func GenerateSlug(value string) string {
	// Convert the string to lowercase
	slug := strings.ToLower(value)

	// Replace spaces with hypens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove any non-alphanumeric characters
	regExp := regexp.MustCompile("[^a-z0-9-]")
	slug = regExp.ReplaceAllString(slug, "")

	// Remove consecutive hypens
	slug = strings.ReplaceAll(slug, "--", "-")

	// Trim leading and trailing hypens
	slug = strings.Trim(slug, "-")

	return slug
}

// ConnectionURLBuilder builds the connection string for the specfic tool
func ConnectionURLBuilder(serviceType string) (string, error) {
	// Define default string
	var connectionString string

	switch serviceType {
	case "postgres":
		// URL for PostgreSQL connection.
		connectionString = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL_MODE"),
		)
		break
	case "redis":
		// URL for Redis connection.
		connectionString = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		// URL for Fiber connection.
		connectionString = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", serviceType)

	}
	return connectionString, nil
}

// CleanString cleans and prepares a string for search.
func CleanString(value string) string {
	if value == "" {
		return ""
	}
	// Trim leading and trailing spaces
	cleaned := strings.TrimSpace(value)

	// Convert to lowercase for case-insensitive search
	cleaned = strings.ToLower(cleaned)

	// Remove symbols and special characters
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleaned = reg.ReplaceAllString(cleaned, "")

	return cleaned
}

// GetPageResponse prepare the page data for the paginated endpoints
func GetPageResponse(totalItems int, page int, limit int) (common.PaginatedResult, error) {
	totalPages := int(totalItems) / limit
	if int(totalItems)%limit != 0 {
		totalPages++
	}

	// Calculate previous and next values
	hasPrevious := page > 1
	hasNext := page < totalPages

	return common.PaginatedResult{
		Page:        page,
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		HasPrevious: hasPrevious,
		HasNext:     hasNext,
	}, nil
}

// GetErrorMessage get the error message for the endpoint
func GetErrorMessage(err error, defaultMessage string) string {
	errorMessage := err.Error()
	if errorMessage == "" {
		errorMessage = defaultMessage
	}
	return errorMessage
}
