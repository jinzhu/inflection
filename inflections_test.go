package inflection

import (
	"strings"
	"testing"
)

var inflections = map[string]string{
	"star":        "stars",
	"STAR":        "STARS",
	"STaR":        "STaRS",
	"Star":        "Stars",
	"bus":         "buses",
	"fish":        "fish",
	"mouse":       "mice",
	"query":       "queries",
	"ability":     "abilities",
	"agency":      "agencies",
	"movie":       "movies",
	"archive":     "archives",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"move":        "moves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"diagnosis":   "diagnoses",
	"diagnosis_a": "diagnosis_as",
	"datum":       "data",
	"medium":      "media",
	"stadium":     "stadia",
	"analysis":    "analyses",
	"node_child":  "node_children",
	"child":       "children",
	"experience":  "experiences",
	"day":         "days",
	"comment":     "comments",
	"foobar":      "foobars",
	"newsletter":  "newsletters",
	"old_news":    "old_news",
	"news":        "news",
	"series":      "series",
	"species":     "species",
	"quiz":        "quizzes",
	"perspective": "perspectives",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"tomato":      "tomatoes",
	"dwarf":       "dwarves",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
	"criterion":   "criteria",
}

func init() {
	AddIrregular("criterion", "criteria")
}

func TestPlural(t *testing.T) {
	for key, value := range inflections {
		if v := Plural(strings.ToUpper(key)); v != strings.ToUpper(value) {
			t.Errorf("%v's plural should be %v, but got %v", strings.ToUpper(key), strings.ToUpper(value), v)
		}

		if v := Plural(strings.Title(key)); v != strings.Title(value) {
			t.Errorf("%v's plural should be %v, but got %v", strings.Title(key), strings.Title(value), v)
		}

		if v := Plural(key); v != value {
			t.Errorf("%v's plural should be %v, but got %v", key, value, v)
		}
	}
}

func TestSingular(t *testing.T) {
	for key, value := range inflections {
		if v := Singular(strings.ToUpper(value)); v != strings.ToUpper(key) {
			t.Errorf("%v's singular should be %v, but got %v", strings.ToUpper(value), strings.ToUpper(key), v)
		}

		if v := Singular(strings.Title(value)); v != strings.Title(key) {
			t.Errorf("%v's singular should be %v, but got %v", strings.Title(value), strings.Title(key), v)
		}

		if v := Singular(value); v != key {
			t.Errorf("%v's singular should be %v, but got %v", value, key, v)
		}
	}
}

func TestGetPlural(t *testing.T) {
	plurals := GetPlural()
	if len(plurals) != len(pluralInflections) {
		t.Errorf("Expected len %d, got %d", len(plurals), len(pluralInflections))
	}
}

func TestGetSingular(t *testing.T) {
	singular := GetSingular()
	if len(singular) != len(singularInflections) {
		t.Errorf("Expected len %d, got %d", len(singular), len(singularInflections))
	}
}

func TestGetIrregular(t *testing.T) {
	irregular := GetIrregular()
	if len(irregular) != len(irregularInflections) {
		t.Errorf("Expected len %d, got %d", len(irregular), len(irregularInflections))
	}
}

func TestGetUncountable(t *testing.T) {
	uncountables := GetUncountable()
	if len(uncountables) != len(uncountableInflections) {
		t.Errorf("Expected len %d, got %d", len(uncountables), len(uncountableInflections))
	}
}

func TestSetPlural(t *testing.T) {
	SetPlural(RegularSlice{{}, {}})
	if len(pluralInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(pluralInflections))
	}
}

func TestSetSingular(t *testing.T) {
	SetSingular(RegularSlice{{}, {}})
	if len(singularInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(singularInflections))
	}
}

func TestSetIrregular(t *testing.T) {
	SetIrregular(IrregularSlice{{}, {}})
	if len(irregularInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(irregularInflections))
	}
}

func TestSetUncountable(t *testing.T) {
	SetUncountable([]string{"", ""})
	if len(uncountableInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(uncountableInflections))
	}
}
