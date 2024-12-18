# Take Home Assessment - Backend Engineering Apprenticeship

Welcome to my submission for the receipt processor challenge! This repository contains the implementation of the solution, along with tests to validate its functionality.

## Requirements

- [Go](https://go.dev/)

## Running the Program

To run the REST API:

```bash
go run .
```
## Testing the Program

To test the program, use the `go test` command:

```bash
go test
```

This will execute all test cases to ensure the correctness of the implementation.

## Project Structure

- `main.go`: Entry point for the application.
- `types.go`: JSON types used for the API.
- `points.go`: Contains logic for point calculation.
- `points_test.go`: Contains test cases for validating the implementation.
- `api.go`: Contains logic for handling HTTP requests.

## Things to Improve

- Receipt persistance.
Currently all receipt calculations are stored in memory. A database would solve this issue.
- A more robust logging system instead of using console output.
- More robust data validation.
Ensure that the incoming HTTP request's data matches the regex pattern provided and ensure date and time fields are valid.