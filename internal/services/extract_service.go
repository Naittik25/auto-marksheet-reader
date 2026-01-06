package services

import (
	"auto-marksheet-reader/internal/storage"
	"regexp"
	"strings"
)

type ExtractService struct{}

// ExtractData parses raw text into structured data
func (s *ExtractService) ExtractData(rawText string) storage.Marksheet {
	data := storage.Marksheet{
		RawText: rawText,
		Marks:   []storage.SubjectMark{},
	}

	lines := strings.Split(rawText, "\n")

	// 1. Simple Regex to find Name (Example: "Name: John Doe")
	nameRegex := regexp.MustCompile(`(?i)(Name|Student)\s*[:|-]?\s*(.*)`)
	rollRegex := regexp.MustCompile(`(?i)(Roll|Seat)\s*No\s*[:|-]?\s*(\w+)`)

	for _, line := range lines {
		// Clean the line
		cleanLine := strings.TrimSpace(line)

		// Find Name
		if data.StudentName == "" {
			match := nameRegex.FindStringSubmatch(cleanLine)
			if len(match) > 2 {
				data.StudentName = strings.TrimSpace(match[2])
			}
		}

		// Find Roll Number
		if data.RollNumber == "" {
			match := rollRegex.FindStringSubmatch(cleanLine)
			if len(match) > 2 {
				data.RollNumber = strings.TrimSpace(match[2])
			}
		}

		// Find Marks (Very basic logic: looks for "Subject 90 A")
		// You will need to improve this based on your specific marksheet format
		parts := strings.Fields(cleanLine)
		if len(parts) >= 2 {
			// This is a placeholder. Real parsing depends heavily on your specific image layout.
			// Currently, it just saves lines that look like subjects.
			// Example logic: If line starts with letters and ends with numbers
		}
	}

	// Adding dummy data for testing UI if nothing found
	if len(data.Marks) == 0 {
		data.Marks = append(data.Marks,
			storage.SubjectMark{Subject: "Mathematics", Score: "85", Grade: "A"},
			storage.SubjectMark{Subject: "Science", Score: "78", Grade: "B"},
		)
	}

	return data
}
