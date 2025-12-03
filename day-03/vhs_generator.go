package main

import (
	"fmt"
	"os"
	"strings"
)

// GenerateVHSTape creates a VHS tape file and a standalone program to record the explanation as a GIF
func GenerateVHSTape(batteries []int, k int, outputFile string) error {
	// Create standalone program
	programFile := "vhs_explanation.go"
	err := createExplanationProgram(batteries, k, programFile)
	if err != nil {
		return fmt.Errorf("failed to create program file: %w", err)
	}

	// Create VHS tape file
	tape := strings.Builder{}
	tape.WriteString("Output " + outputFile + "\n")
	tape.WriteString("Set FontSize 12\n")
	tape.WriteString("Set Width 720\n")
	tape.WriteString("Set Height 480\n")
	tape.WriteString("Set Theme \"Molokai\"\n")
	tape.WriteString("Set Shell \"bash\"\n")
	tape.WriteString("\n")
	tape.WriteString(fmt.Sprintf("Type \"go run %s super_joltage_explained.go\"\n", programFile))
	tape.WriteString("Sleep 1s\n")
	tape.WriteString("Enter\n")
	tape.WriteString("Sleep 140s\n")

	tapeFile := "explanation.tape"
	err = os.WriteFile(tapeFile, []byte(tape.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write tape file: %w", err)
	}

	fmt.Printf("Created files:\n")
	fmt.Printf("  - %s (VHS tape file)\n", tapeFile)
	fmt.Printf("  - %s (standalone program)\n", programFile)
	fmt.Printf("\nTo generate GIF, run:\n")
	fmt.Printf("  vhs %s\n", tapeFile)
	fmt.Printf("\nThis will create: %s\n", outputFile)

	return nil
}

func createExplanationProgram(batteries []int, k int, filename string) error {
	content := fmt.Sprintf(`//go:build ignore

package main

func main() {
	batteries := []int{%s}
	SuperJoltageExplained(batteries, %d)
}
`, formatBatteries(batteries), k)

	return os.WriteFile(filename, []byte(content), 0644)
}

func formatBatteries(batteries []int) string {
	parts := make([]string, len(batteries))
	for i, b := range batteries {
		parts[i] = fmt.Sprintf("%d", b)
	}
	return strings.Join(parts, ", ")
}
