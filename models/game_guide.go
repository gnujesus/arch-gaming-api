package models

import "time"

type GameGuide struct {
	ID                  int
	GameID              int    // fk
	CompatibilityRating string // perfect, excellent, pretty darn good, regular, mediocre, bad, doesn't run at all
	SetupInstructions   string // markdown text
	RequiredPackages    string
	KnownIssues         string
	TestedAt            time.Time
}
