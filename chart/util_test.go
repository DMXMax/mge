package chart

import (
	"testing"
)

// write a test that test RollOdds
// hint: use the testing package
// hint: use the RollOdds function
func TestRollOdds(t *testing.T) {
	t.Log("Testing RollOdds")
	for i := 1; i <= 9; i++ {
		res := FateChart.RollOdds(VeryLikely, i)
		t.Logf("%#v\n", res)
	}

}

func TestEvaluate(t *testing.T) {
	t.Log("Testing Evaluate")
	res := evaluate(5, 82)
	want := "Exceptional No"
	if res.Text != "Exceptional No" {
		t.Errorf("Expected %s , got %s", want, res.Text)
	}
	if res = evaluate(5, 81); res.Text != "No" {
		t.Errorf("Expected No, got %s", res.Text)
	}
	if res = evaluate(-20, 77); res.Text != "Exceptional No" {
		t.Errorf("Expected No, got %s", res.Text)
	}
	if res = evaluate(50, 91); res.Text != "Exceptional No" {
		t.Errorf("Expected No, got %s", res.Text)
	}
	if res = evaluate(50, 10); res.Text != "Exceptional Yes" {
		t.Errorf("Expected Exceptional Yes, got %s", res.Text)
	}
	if res = evaluate(50, 11); res.Text != "Yes" {
		t.Errorf("Expected Yes, got %s", res.Text)
	}
}
