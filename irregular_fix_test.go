package inflection

import "testing"

func TestHumanAndQuota(t *testing.T) {
	if got := Plural("human"); got != "humans" {
		t.Fatalf("Plural(human)=%q", got)
	}
	if got := Singular("humans"); got != "human" {
		t.Fatalf("Singular(humans)=%q", got)
	}
	if got := Plural("Human"); got != "Humans" {
		t.Fatalf("Plural(Human)=%q", got)
	}
	if got := Plural("man"); got != "men" {
		t.Fatalf("Plural(man)=%q", got)
	}
	if got := Plural("quota"); got != "quotas" {
		t.Fatalf("Plural(quota)=%q", got)
	}
	if got := Singular("quotas"); got != "quota" {
		t.Fatalf("Singular(quotas)=%q", got)
	}
	if got := Singular("quota"); got != "quota" {
		t.Fatalf("Singular(quota)=%q want quota (not quotum)", got)
	}
}
