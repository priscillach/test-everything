package b2t

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"testing"
	"time"
)

func TestConvertBookmarksToRules(t *testing.T) {
	rulesOutput, err := convertBookmarksToRules("your chrome bookmarks html path")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonOutput, err := json.MarshalIndent(rulesOutput, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	currentTime := time.Now()
	randomDigits := fmt.Sprintf("%06d", rand.Intn(1000000))
	filename := fmt.Sprintf("tabgroups_rules_%s_%s.json", currentTime.Format("20060102"), randomDigits)

	outputDir := "your dir path to save the tab groups rule json" // Replace this with your desired directory path
	fullPath := filepath.Join(outputDir, filename)

	err = ioutil.WriteFile(fullPath, jsonOutput, 0644)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

}
