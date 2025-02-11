package themes

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentTheme() (string, error) {
	configPath := os.Getenv("HYDE_STATE_HOME") + "/staterc"
	file, err := os.Open(configPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "HYDE_THEME=") {
			return strings.Trim(strings.Split(line, "=")[1], `"`), nil
		}
	}

	return "", fmt.Errorf("HYDE_THEME не найдена")
}

var themesDir = os.Getenv("XDG_CONFIG_HOME") + "/hyde/themes"

func GetThemes() ([]string, error) {
	files, err := os.ReadDir(themesDir)
	if err != nil {
		return nil, err
	}

	var themes []string
	for _, file := range files {
		if file.IsDir() {
			themes = append(themes, file.Name())
		} else if filepath.Ext(file.Name()) == ".json" || filepath.Ext(file.Name()) == ".theme" {
			themes = append(themes, file.Name())
		}
	}

	if len(themes) == 0 {
		return nil, fmt.Errorf("темы не найдены в %s", themesDir)
	}

	return themes, nil
}

func SetTheme(newTheme string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("source ~/.local/lib/hyde/globalcontrol.sh && ~/.local/lib/hyde/themeswitch.sh -qs \"%s\"", newTheme))
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
