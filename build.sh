#!/bin/bash

# Set Swagger output directory
SWAGGER_OUTPUT="docs"

# Function to generate Swagger docs
generate_swagger_docs() {
    echo "Generating Swagger docs..."
    swag fmt
    # https://github.com/swaggo/swag/issues/817
    swag init --parseDependency -g cmd/main.go --output $SWAGGER_OUTPUT
}

# Function to run the application
run_app() {
    echo "Running the application..."
    go run cmd/main.go
}

# Main function to handle swag and run targets
main() {
    case "$1" in
        swag)
            generate_swagger_docs
            ;;
        run)
            run_app
            ;;
        app)
            generate_swagger_docs
            run_app
            ;;
        *)
            echo "Usage: $0 {swag|run|app}"
            exit 1
    esac
}

# Call the main function with the first argument
main "$1"
