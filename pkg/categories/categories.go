// Package categories returns a list of categories
package categories

// Categories returns a simple map of category codes against their
// full category names.
// Category codes are used in Shoot 'n Score It competitor entries,
// while the full category name is required for Practiscore import
func Categories() map[string]string {
	m := map[string]string{
		"J":  "junior",
		"SJ": "super junior",
		"L":  "lady",
		"S":  "senior",
		"SS": "super senior",
	}
	return m
}
