package util

import (
	"testing"
)

func TestGetEventFocus(t *testing.T) {
	for i := 0; i < 100; i++ {
		focus := GetEventFocus()
		if _, ok := EventText[focus]; !ok {
			t.Fatalf("GetEventFocus() returned an invalid EventFocus value: %d", focus)
		}
	}
}

func TestGetEventAction(t *testing.T) {
	action, subject := GetEventAction()

	if action == "" {
		t.Error("GetEventAction() returned an empty action")
	}
	if subject == "" {
		t.Error("GetEventAction() returned an empty subject")
	}

	// Helper to check if a slice contains a string
	contains := func(slice []string, val string) bool {
		for _, item := range slice {
			if item == val {
				return true
			}
		}
		return false
	}

	if !contains(Action, action) {
		t.Errorf("returned action %q not found in Action table", action)
	}
	if !contains(Subject, subject) {
		t.Errorf("returned subject %q not found in Subject table", subject)
	}
}

func TestGetMeaningActions(t *testing.T) {
	actions := GetMeaningActions()
	if len(actions) != 2 {
		t.Fatalf("expected 2 meaning actions, got %d", len(actions))
	}
	if actions[0] == "" || actions[1] == "" {
		t.Error("got an empty meaning action string")
	}
}

func TestGetMeaningDescriptors(t *testing.T) {
	descriptors := GetMeaningDescriptors()
	if len(descriptors) != 2 {
		t.Fatalf("expected 2 meaning descriptors, got %d", len(descriptors))
	}
	if descriptors[0] == "" || descriptors[1] == "" {
		t.Error("got an empty meaning descriptor string")
	}
}

func TestGetEvent(t *testing.T) {
	event := GetEvent()
	if event == nil {
		t.Fatal("GetEvent() returned nil")
	}

	if _, ok := EventText[event.Focus]; !ok {
		t.Errorf("event has invalid focus: %d", event.Focus)
	}
	if event.Action == "" {
		t.Error("event action is empty")
	}
	if event.Subject == "" {
		t.Error("event subject is empty")
	}
	if len(event.Meaning.Actions) != 2 || event.Meaning.Actions[0] == "" || event.Meaning.Actions[1] == "" {
		t.Errorf("event meaning actions are invalid: %v", event.Meaning.Actions)
	}
	if len(event.Meaning.Descriptors) != 2 || event.Meaning.Descriptors[0] == "" || event.Meaning.Descriptors[1] == "" {
		t.Errorf("event meaning descriptors are invalid: %v", event.Meaning.Descriptors)
	}
}

func TestEvent_String(t *testing.T) {
	event := &Event{
		Focus:   NPCAction,
		Action:  "Guide",
		Subject: "Power",
		Meaning: struct {
			Actions     []string
			Descriptors []string
		}{Actions: GetMeaningActions(), Descriptors: GetMeaningDescriptors()},
	}
	expected := "NPC Action: Guide Power (Boldly Mighty, Take Advantage)"
	if got := event.String(); got != expected {
		t.Errorf("String() mismatch:\n got: %q\nwant: %q", got, expected)
	}
}
