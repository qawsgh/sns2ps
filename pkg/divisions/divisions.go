// Package divisions returns a simple map of division codes against their
// full division names.
// Division codes are used in Shoot 'n Score It competitor entries,
// while the full division name is required for Practiscore import
package divisions

// Divisions returns the list of divisions to be used to translate codes to names
func Divisions() map[string]string {
	m := map[string]string{
		"hg1":  "Open",
		"hg2":  "Standard",
		"hg3":  "Production",
		"hg4":  "Modified",
		"hg5":  "Revolver",
		"hg6":  "Open",
		"hg7":  "Production",
		"hg8":  "Single-Stack",
		"hg9":  "Limited",
		"hg10": "Limited-10",
		"hg11": "Revolver",
		"hg12": "Classic",
		"hgc":  "Custom",
		"rf1":  "Semi-Auto Open",
		"rf2":  "Semi-Auto Standard",
		"rf3":  "Manual Action Open",
		"rf4":  "Manual Action Standard",
		"rf5":  "Open",
		"rf6":  "Standard",
		"rf7":  "Tactical",
		"rf8":  "Manually Operated",
		"rf9":  "Semi-Auto",
		"rf10": "Manually Operated",
		"rf11": "Manual Action Standard 10",
		"rfc":  "Custom",
		"mr1":  "Mini Rifle Open",
		"mr2":  "Mini Rifle Standard",
		"mrc":  "Mini Rifle Custom",
		"sg1":  "Open",
		"sg2":  "Modified",
		"sg3":  "Standard",
		"sg4":  "Standard Manual",
		"sg5":  "Open",
		"sg6":  "Limited/Tactical",
		"sg7":  "Heavy Metal",
		"sgc":  "Custom",
		"ai1":  "Open",
		"ai2":  "Standard",
		"ai3":  "Production",
		"ai3a": "Production Optics",
		"ai4":  "Open",
		"ai5":  "Standard",
		"ai6":  "Production",
		"ai8":  "Classic",
		"aic":  "Custom",
	}
	return m
}
