package sender
import (
	"encoding/json"
	"fmt"
	"kt_project/internal/models"
	"net/http"
	"bytes"

)
func SendStat(url string, data *models.Stat) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data stat: %w", err)
	}

	// making post request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// checking obj status for 201 (created)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error server response: %s", resp.Status)
	}

	fmt.Printf("[AGENT] successed sending data. status: %s\n", resp.Status)
	return nil
}
