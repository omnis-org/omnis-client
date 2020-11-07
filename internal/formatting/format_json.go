package formatting

import (
	"encoding/json"
	"fmt"
	"omnis-client/core"
	"os"
	"time"
)

// ExportJSON will save the output to a JSON file
func ExportJSON(out core.Output) error {

	jsonstr, err := json.Marshal(out)
	if err != nil {
		return err
	}

	now := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("./omnis_%s.json", now)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(jsonstr)
	if err != nil {
		return err
	}

	fmt.Printf("Output as json : %s", filename)

	return nil
}
