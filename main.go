package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting the performance auditor...")

	// Check if the required environment variable is set
	if os.Getenv("PERFORMANCE_AUDIT_ENABLED") == "" {
		fmt.Println("PERFORMANCE_AUDIT_ENABLED is not set. Exiting...")
		return
	}
	
	lookback_days := os.Getenv("INPUT_LOOKBACK_DAYS")
	if lookback_days == "" {
		fmt.Println("INPUT_LOOKBACK_DAYS is not set. Exiting...")
		return
	}

	output_file := os.Getenv("GITHUB_OUTPUT_FILE")
	report_url := fmt.Sprintf("report-url=https://example.com/%s", output_file)

	err := os.WriteFile(output_file, []byte(report_url), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file %s: %v\n", output_file, err)
		return
	}

	fmt.Println("Report generated and output set successfully.")
}