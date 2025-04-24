package parser

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"

	model "lexigo/model/china"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var cedictPattern = regexp.MustCompile(`^(\S+)\s+(\S+)\s+\[(.+?)\]\s+/(.+)/$`)

func ParseCCCEDICT(path string, db *gorm.DB) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	successCount := 0
	errorCount := 0
	skipCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Bỏ qua dòng comment
		if strings.HasPrefix(line, "#") {
			skipCount++
			continue
		}

		matches := cedictPattern.FindStringSubmatch(line)
		if len(matches) != 5 {
			skipCount++
			continue // Không khớp định dạng
		}

		entry := model.ChinaEntry{
			Base: model.Base{
				ID: uuid.New(),
			},
			Traditional: &matches[1],
			Simplified:  &matches[2],
			Pinyin:      &matches[3],
			Definition:  &matches[4],
		}

		if err := db.Create(&entry).Error; err != nil {
			errorCount++
			// Log lỗi nhưng tiếp tục xử lý các entry khác
			log.Printf("Error inserting entry: %v, Line: %s", err, line)
			continue
		}
		successCount++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	log.Printf("Parsing completed. Success: %d, Errors: %d, Skipped: %d", successCount, errorCount, skipCount)
	return nil
}

func CountCCCEDICT(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Bỏ qua dòng comment
		if strings.HasPrefix(line, "#") {
			continue
		}

		matches := cedictPattern.FindStringSubmatch(line)
		if len(matches) == 5 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
