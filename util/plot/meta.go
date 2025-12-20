package plot

import "fmt"

// MetaPlotPoint represents a single entry in the meta plot points table.
type MetaPlotPoint struct {
	// Upper bound of the dice roll for this entry (inclusive).
	Range int
	// Text is the descriptive text for the meta plot point.
	Text string
}

// MetaPlotPointsTable contains the structured data for meta plot points.
// The table is ordered by the roll range.
var MetaPlotPointsTable = []MetaPlotPoint{
	{
		Range: 18,
		Text: `CHARACTER EXITS THE ADVENTURE: A Character, who is not a Player Character, is removed from the
Characters List completely. Cross out all references to that Character on the Characters List. If there are no non-
Player Characters, then re-roll for another Meta Plot Point. This change can be reflected in the activity in this Turning
Point or not. For instance, you may explain the Character being removed from the Adventure by having that Character
die in the Turning Point. Or, you simply remove them from the Characters List and decide that their involvement in
the Adventure is over. If, when rolling on the Characters List to determine who this Character is, you roll a Player
Character or “New Character”, then consider it a result of “Choose The Most Logical Character”.`,
	},
	{
		Range: 27,
		Text: `CHARACTER RETURNS: A Character who previously had been removed from the Adventure returns. Write that
Character back into the Characters List with a single listing. If there are no Characters to return, then treat this as a
“New Character” result and use this Plot Point to introduce a new Character into the Turning Point. If there is more
than one Character who can return, then choose the most logical Character to return. This change can be reflected
in the activity in this Turning Point or not.`,
	},
	{
		Range: 36,
		Text: `CHARACTER STEPS UP: A Character becomes more important, gaining another slot on the Characters List
even if it pushes them past 3 slots. When you roll on the Characters List to see who the Character is, treat a
result of “New Character” as “Choose The Most Logical Character”. This change can be reflected in the activity in
this Turning Point or not.`,
	},
	{
		Range: 55,
		Text: `CHARACTER STEPS DOWN: A Character becomes less important, remove them from one slot on the Characters
List even if it removes them completely from the List. If this would remove a Player Character completely from the List,
or if when rolling for the Character you get a result of “New Character”, then treat this as a result of “Choose The Most
Logical Character”. If there is no possible Character to choose without removing a Player Character completely from the
List, then roll again on the Meta Plot Points Table. This change can be reflected in the activity in this Turning Point or not.`,
	},
	{
		Range: 73,
		Text: `CHARACTER DOWNGRADE: A Character becomes less important, remove them from two slots on the Characters
List even if it removes them completely from the List. If this would remove a Player Character completely from the List,
or if when rolling for the Character you get a result of “New Character”, then treat this as a result of “Choose The Most
Logical Character”. If there is no possible Character to choose without removing a Player Character completely from the
List, then roll again on the Meta Plot Points Table. This change can be reflected in the activity in this Turning Point or not.`,
	},
	{
		Range: 82,
		Text: `CHARACTER UPGRADE: A Character becomes more important, gaining 2 slots on the Characters List even
if it pushes them past 3 slots. When you roll on the Characters List to see who the Character is, treat a result
of “New Character” as “Choose The Most Logical Character”. This change can be reflected in the activity in this
Turning Point or not.`,
	},
	{
		Range: 100,
		Text: `PLOTLINE COMBO: This Turning Point is about more than one Plotline at the same time. Roll again on the
Plotlines List and add that Plotline to this Turning Point along with the original Plotline rolled. If when rolling for
an additional Plotline you roll the same Plotline already in use for this Turning Point, then treat the result as a
“Choose The Most Logical Plotline”. If there are no other Plotlines to choose from, then create a new Plotline as
the additional Plotline. If a Conclusion is rolled as a Plot Point during this Turning Point, apply it to the Plotline
that seems most appropriate. If another Conclusion is rolled, continue to apply them to the additional Plotlines
in this Turning Point if you can. It is possible with repeated results of “Plotline Combo” to have more than two
Plotlines combined in this way.`,
	},
}

// GetMetaPlotPoint finds the corresponding meta plot point for a given d100 roll.
func GetMetaPlotPoint(roll int) (*MetaPlotPoint, error) {
	if roll < 1 || roll > 100 {
		return nil, fmt.Errorf("roll out of range (1-100): %d", roll)
	}

	for i := range MetaPlotPointsTable {
		point := &MetaPlotPointsTable[i]
		if roll <= point.Range {
			return point, nil
		}
	}

	// This should be unreachable if the table is correctly defined up to 100.
	return nil, fmt.Errorf("no meta plot point found for roll %d", roll)
}
