package band

import "testing"

func TestBandwidth(t *testing.T) {
	for _, item := range BandListAll() {
		if item.Bandwidth() != item.UpperFreqMHz-item.LowerFreqMHz {
			t.Errorf("Band %s bandwidth is incorrect: %f", item.Key, item.Bandwidth())
		}
	}
}

func TestIsInBand(t *testing.T) {
	for _, item := range BandListAll() {
		if !item.IsInBand(item.LowerFreqMHz) {
			t.Errorf("Band %s should be in band", item.Key)
		}
		if !item.IsInBand(item.UpperFreqMHz) {
			t.Errorf("Band %s should be in band", item.Key)
		}
		if item.IsInBand(item.LowerFreqMHz - 1) {
			t.Errorf("Band %s should not be in band", item.Key)
		}
		if item.IsInBand(item.UpperFreqMHz + 1) {
			t.Errorf("Band %s should not be in band", item.Key)
		}
	}
}

func TestFindBandByMHz(t *testing.T) {
	for _, item := range BandListAll() {
		_, ok := FindBandByMHz(item.LowerFreqMHz)
		if !ok {
			t.Errorf("Band %s should be in band", item.Key)
		}
	}
}

func TestFindBandByMHz_NotInBand(t *testing.T) {
	for _, item := range BandListAll() {
		_, ok := FindBandByMHz(0.25)
		if ok {
			t.Errorf("Band %s should be in band", item.Key)
		}
	}
}
