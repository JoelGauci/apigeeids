package feodotracker

import "testing"

func TestGetRules(t *testing.T) {
	var feodotracker_url string = "https://feodotracker.abuse.ch/downloads/feodotracker.rules"
	r, err := New(feodotracker_url, nil, "FeodoTracker test GetRules()")
	if err != nil {
		t.Errorf("Ruler of type FeodoTracker not created (Ruler=%v)", r)
	}

	err = r.GetRules()
	if err != nil {
		t.Errorf("Ruler FeodoTracker cannot get rules (Ruler=%v - Error:%v)", r, err)
	}

	if len(r.Lines) <= 1 {
		t.Errorf("Rules have not been retrieved correctly (number of lines:%d)", len(r.Lines))
	}
}

func TestSetRules(t *testing.T) {
	var feodotracker_url string = "https://feodotracker.abuse.ch/downloads/feodotracker.rules"
	r, err := New(feodotracker_url, nil, "FeodoTracker test GetRules()")
	if err != nil {
		t.Errorf("Ruler of type FeodoTracker not created (Ruler=%v)", r)
	}

	err = r.GetRules()
	if err != nil {
		t.Errorf("Ruler FeodoTracker cannot get rules (Ruler=%v - Error:%v)", r, err)
	}

	r.SetRules()
	if len(r.Rules) <= 1 { // at least 2 lines. Current date is inserted as a banner + at least one following rule !
		t.Errorf("Rules have not been set correctly (number of rules:%d)", len(r.Rules))
	}

}
