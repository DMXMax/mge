package util

import (
	"math/rand"

	"github.com/DMXMax/mge/util/elements"
)

type EventFocus int

const (
	Remote EventFocus = iota
	Ambiguous
	NewNPC
	NPCAction
	NPCNegative
	NPCPositive
	MoveTowardThread
	MoveAwayFromThread
	CloseThread
	PCNegative
	PCPositive
	CurrentContext
)

var EventText = map[EventFocus]string{
	Remote:             "Remote event",
	Ambiguous:          "Ambiguous event",
	NewNPC:             "New NPC",
	NPCAction:          "NPC Action",
	NPCNegative:        "NPC Negative",
	NPCPositive:        "NPC Positive",
	MoveTowardThread:   "Move Toward a Thread",
	MoveAwayFromThread: "Move Away From a Thread",
	CloseThread:        "Close a Thread",
	PCNegative:         "PC Negative",
	PCPositive:         "PC Positive",
	CurrentContext:     "Current Context",
}

/*
1-7 . .
8-28 29-35 36-45 46-52 . 53-55 56-67 68-75 76-83 84-92 93-100
. . . . . . . . . .
Event Focus Table
. . . . . . . . .Remote event
. . . . . . . . .NPC action
. . . . . . . . .Introduce a new NPC
. . . . . . . . .Move toward a thread
. . . . . . . . .Move away from a thread . . . . . . . . .Close a thread
. . . . . . . . .PC negative
. . . . . . . . .PC positive
. . . . . . . . .Ambiguous event
. . . . . . . . .NPC negative
. . . . . . . . .NPC positive
*/
//randon number from 1 to 100
func GetEventFocus() EventFocus {
	//random number from 1 to 100

	switch roll := rand.Intn(100) + 1; {
	case roll <= 5:
		return Remote
	case roll <= 10:
		return Ambiguous
	case roll <= 20:
		return NewNPC
	case roll <= 40:
		return NPCAction
	case roll <= 45:
		return NPCNegative
	case roll <= 50:
		return NPCPositive
	case roll <= 55:
		return MoveTowardThread
	case roll <= 65:
		return MoveAwayFromThread
	case roll <= 70:
		return CloseThread
	case roll <= 80:
		return PCNegative
	case roll <= 85:
		return PCPositive
	default:
		return CurrentContext

	}
}

func GetEventAction() (string, string) {

	return Action[rand.Intn(len(Action))], Subject[rand.Intn(len(Subject))]

}

func GetMeaningActions() []string {
	return []string{
		elements.ActionTable1[rand.Intn(len(elements.ActionTable1))],
		elements.ActionTable2[rand.Intn(len(elements.ActionTable2))],
	}
}

func GetMeaningDescriptors() []string {
	return []string{
		elements.Descriptor1[rand.Intn(len(elements.Descriptor1))],
		elements.Descriptor2[rand.Intn(len(elements.Descriptor2))],
	}
}

type Event struct {
	Focus           EventFocus
	Action, Subject string
	Meaning         struct {
		Actions     []string
		Descriptors []string
	}
}

func (e Event) String() string {
	return EventText[e.Focus] + ": " + e.Action + " " + e.Subject + " (" + e.Meaning.Descriptors[0] + " " + e.Meaning.Descriptors[1] + ", " + e.Meaning.Actions[0] + " " + e.Meaning.Actions[1] + ")"
}

func GetEvent() *Event {
	focus := GetEventFocus()
	action, subject := GetEventAction()
	return &Event{focus, action, subject, struct {
		Actions     []string
		Descriptors []string
	}{Actions: GetMeaningActions(), Descriptors: GetMeaningDescriptors()}}
}
