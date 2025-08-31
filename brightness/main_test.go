package brightness

import "testing"

func TestChangeBrightness(t *testing.T) {
	err := ChangeBrightness(10)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
}
