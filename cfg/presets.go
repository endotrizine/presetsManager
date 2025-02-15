package cfg

import (
	"encoding/json"
	"fmt"
	"os"
)

// Preset - структура одного пресета
type Preset struct {
	Name     string            `json:"name"`
	Selected bool              `json:"selected"`
	Params   map[string]string `json:"params"`
}

// LoadPresets загружает пресеты из JSON-файла
func LoadPresets() ([]Preset, error) {
	filePath := "cfg/presets.json"
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}

	var data struct {
		Presets []Preset `json:"presets"`
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return data.Presets, nil
}
