package migrations

import (
	"fmt"

	"github.com/ahakkila/migrator"
)

var Version int = 1

func Initialize_001(m *migrator.Migrator) {
	m.NewStep(Version, Up_001, Down_001)
}

func Up_001() {
	fmt.Println("   Do migrations for version 001")
}

func Down_001() {
	fmt.Println("   Revert migrations for version 001")
}
