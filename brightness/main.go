package brightness

import (
	"fmt"
	"os"
)

func ChangeBrightness(value uint8) error {
	if value > 100 {
		return fmt.Errorf("value should be between 1 and 100")
	}

	filepath := "/sys/class/backlight/nvidia_wmi_ec_backlight/brightness"

	file, err := os.OpenFile(filepath, os.O_WRONLY, 0660)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	data := []byte{value}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
