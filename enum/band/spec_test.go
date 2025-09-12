package band

import (
	"testing"
)

func TestBandwidth(t *testing.T) {
	for _, item := range ListActive() {
		if item.Bandwidth() != float64(item.UpperFreqMHz)-float64(item.LowerFreqMHz) {
			t.Errorf("Band %s bandwidth is incorrect: %f", item.Key, item.Bandwidth())
		}
	}
}

func TestIsInBand(t *testing.T) {
	for _, item := range ListActive() {
		if !item.IsInBand(float64(item.LowerFreqMHz)) {
			t.Errorf("frequency %f should be in band %s", item.LowerFreqMHz, item.Key)
		}
		if !item.IsInBand(float64(item.UpperFreqMHz)) {
			t.Errorf("frequency %f should be in band %s", item.UpperFreqMHz, item.Key)
		}
		if item.IsInBand(float64(item.LowerFreqMHz) - .01) {
			t.Errorf("frequency %f should NOT be in band %s", item.LowerFreqMHz-.01, item.Key)
		}
		if item.IsInBand(float64(item.UpperFreqMHz) + .01) {
			t.Errorf("frequency %f should NOT be in band %s", item.UpperFreqMHz+.01, item.Key)
		}
	}
}

func TestFindBandByMHz(t *testing.T) {
	b, ok := FindBandByMHz(7.050)
	if !ok || b.Key != BAND_40M {
		t.Errorf("7.050 MHz should be in band %s", BAND_40M)
	}
}

func TestFindBandByMHz_NotInBand(t *testing.T) {
	for _, item := range ListActive() {
		_, ok := FindBandByMHz(0.25)
		if ok {
			t.Errorf("0.25Mhz should not be in band %s", item.Key)
		}
	}
}
