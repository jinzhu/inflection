/*
Package inflection pluralizes and singularizes English nouns.

		inflection.Plural("person") => "people"
		inflection.Plural("Person") => "People"
		inflection.Plural("PERSON") => "PEOPLE"

		inflection.Singularize("people") => "person"
		inflection.Singularize("People") => "Person"
		inflection.Singularize("PEOPLE") => "PERSON"

		inflection.Plural("FancyPerson") => "FancydPeople"
		inflection.Singularize("FancyPeople") => "FancydPerson"

Standard rules are from Rails's ActiveSupport (https://github.com/rails/rails/blob/master/activesupport/lib/active_support/inflections.rb)

If you want to register more rules, follow:

		inflection.AddUncountable("fish")
		inflection.AddIrregular("person", "people")
		inflection.AddPlural("(bu)s$", "${1}ses") # "bus" => "buses" / "BUS" => "BUSES" / "Bus" => "Buses"
		inflection.AddSingular("(bus)(es)?$", "${1}") # "buses" => "bus" / "Buses" => "Bus" / "BUSES" => "BUS"
*/
package inflection

import (
	"regexp"
	"strings"
)

type inflection struct {
	regexp  *regexp.Regexp
	replace string
}

type Regular struct {
	find    string
	replace string
}

type Irregular struct {
	singular string
	plural   string
}

type RegularSlice []Regular
type IrregularSlice []Irregular

var pluralInflections = RegularSlice{
	{find: "([a-z])$", replace: "${1}s"},
	{find: "s$", replace: "s"},
	{find: "^(ax|test)is$", replace: "${1}es"},
	{find: "(octop|vir)us$", replace: "${1}i"},
	{find: "(octop|vir)i$", replace: "${1}i"},
	{find: "(alias|status)$", replace: "${1}es"},
	{find: "(bu)s$", replace: "${1}ses"},
	{find: "(buffal|tomat)o$", replace: "${1}oes"},
	{find: "([ti])um$", replace: "${1}a"},
	{find: "([ti])a$", replace: "${1}a"},
	{find: "sis$", replace: "ses"},
	{find: "(?:([^f])fe|([lr])f)$", replace: "${1}${2}ves"},
	{find: "(hive)$", replace: "${1}s"},
	{find: "([^aeiouy]|qu)y$", replace: "${1}ies"},
	{find: "(x|ch|ss|sh)$", replace: "${1}es"},
	{find: "(matr|vert|ind)(?:ix|ex)$", replace: "${1}ices"},
	{find: "^(m|l)ouse$", replace: "${1}ice"},
	{find: "^(m|l)ice$", replace: "${1}ice"},
	{find: "^(ox)$", replace: "${1}en"},
	{find: "^(oxen)$", replace: "${1}"},
	{find: "(quiz)$", replace: "${1}zes"},
}

var singularInflections = RegularSlice{
	{find: "s$", replace: ""},
	{find: "(ss)$", replace: "${1}"},
	{find: "(n)ews$", replace: "${1}ews"},
	{find: "([ti])a$", replace: "${1}um"},
	{find: "((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$", replace: "${1}sis"},
	{find: "(^analy)(sis|ses)$", replace: "${1}sis"},
	{find: "([^f])ves$", replace: "${1}fe"},
	{find: "(hive)s$", replace: "${1}"},
	{find: "(tive)s$", replace: "${1}"},
	{find: "([lr])ves$", replace: "${1}f"},
	{find: "([^aeiouy]|qu)ies$", replace: "${1}y"},
	{find: "(s)eries$", replace: "${1}eries"},
	{find: "(m)ovies$", replace: "${1}ovie"},
	{find: "(c)ookies$", replace: "${1}ookie"},
	{find: "(x|ch|ss|sh)es$", replace: "${1}"},
	{find: "^(m|l)ice$", replace: "${1}ouse"},
	{find: "(bus)(es)?$", replace: "${1}"},
	{find: "(o)es$", replace: "${1}"},
	{find: "(shoe)s$", replace: "${1}"},
	{find: "(cris|test)(is|es)$", replace: "${1}is"},
	{find: "^(a)x[ie]s$", replace: "${1}xis"},
	{find: "(octop|vir)(us|i)$", replace: "${1}us"},
	{find: "(alias|status)(es)?$", replace: "${1}"},
	{find: "^(ox)en", replace: "${1}"},
	{find: "(vert|ind)ices$", replace: "${1}ex"},
	{find: "(matr)ices$", replace: "${1}ix"},
	{find: "(quiz)zes$", replace: "${1}"},
	{find: "(database)s$", replace: "${1}"},
}

var irregularInflections = IrregularSlice{
	{singular: "person", plural: "people"},
	{singular: "man", plural: "men"},
	{singular: "child", plural: "children"},
	{singular: "sex", plural: "sexes"},
	{singular: "move", plural: "moves"},
	{singular: "mombie", plural: "mombies"},
}

var uncountableInflections = []string{"equipment", "information", "rice", "money", "species", "series", "fish", "sheep", "jeans", "police"}

var compiledPluralMaps []inflection
var compiledSingularMaps []inflection

func compile() {
	compiledPluralMaps = []inflection{}
	compiledSingularMaps = []inflection{}
	for _, uncountable := range uncountableInflections {
		inf := inflection{
			regexp:  regexp.MustCompile("^(?i)(" + uncountable + ")$"),
			replace: "${1}",
		}
		compiledPluralMaps = append(compiledPluralMaps, inf)
		compiledSingularMaps = append(compiledSingularMaps, inf)
	}

	for _, value := range irregularInflections {
		infs := []inflection{
			inflection{regexp: regexp.MustCompile(strings.ToUpper(value.singular) + "$"), replace: strings.ToUpper(value.plural)},
			inflection{regexp: regexp.MustCompile(strings.Title(value.singular) + "$"), replace: strings.Title(value.plural)},
			inflection{regexp: regexp.MustCompile(value.singular + "$"), replace: value.plural},
		}
		compiledPluralMaps = append(compiledPluralMaps, infs...)
	}

	for _, value := range irregularInflections {
		infs := []inflection{
			inflection{regexp: regexp.MustCompile(strings.ToUpper(value.plural) + "$"), replace: strings.ToUpper(value.singular)},
			inflection{regexp: regexp.MustCompile(strings.Title(value.plural) + "$"), replace: strings.Title(value.singular)},
			inflection{regexp: regexp.MustCompile(value.plural + "$"), replace: value.singular},
		}
		compiledSingularMaps = append(compiledSingularMaps, infs...)
	}

	for i := len(pluralInflections) - 1; i >= 0; i-- {
		value := pluralInflections[i]
		infs := []inflection{
			inflection{regexp: regexp.MustCompile(strings.ToUpper(value.find)), replace: strings.ToUpper(value.replace)},
			inflection{regexp: regexp.MustCompile(value.find), replace: value.replace},
			inflection{regexp: regexp.MustCompile("(?i)" + value.find), replace: value.replace},
		}
		compiledPluralMaps = append(compiledPluralMaps, infs...)
	}

	for i := len(singularInflections) - 1; i >= 0; i-- {
		value := singularInflections[i]
		infs := []inflection{
			inflection{regexp: regexp.MustCompile(strings.ToUpper(value.find)), replace: strings.ToUpper(value.replace)},
			inflection{regexp: regexp.MustCompile(value.find), replace: value.replace},
			inflection{regexp: regexp.MustCompile("(?i)" + value.find), replace: value.replace},
		}
		compiledSingularMaps = append(compiledSingularMaps, infs...)
	}
}

func init() {
	compile()
}

func AddPlural(find, replace string) {
	pluralInflections = append(pluralInflections, Regular{find: find, replace: replace})
	compile()
}

func AddSingular(find, replace string) {
	singularInflections = append(singularInflections, Regular{find: find, replace: replace})
	compile()
}

func AddIrregular(singular, plural string) {
	irregularInflections = append(irregularInflections, Irregular{singular, plural})
	compile()
}

func AddUncountable(value string) {
	uncountableInflections = append(uncountableInflections, value)
	compile()
}

func GetPlural() RegularSlice {
	plurals := make(RegularSlice, len(pluralInflections))
	copy(plurals, pluralInflections)
	return plurals
}

func GetSingular() RegularSlice {
	singulars := make(RegularSlice, len(singularInflections))
	copy(singulars, singularInflections)
	return singulars
}

func GetIrregular() IrregularSlice {
	irregular := make(IrregularSlice, len(irregularInflections))
	copy(irregular, irregularInflections)
	return irregular
}

func GetUncountable() []string {
	uncountables := make([]string, len(uncountableInflections))
	copy(uncountables, uncountableInflections)
	return uncountables
}

func SetPlural(inflections RegularSlice) {
	pluralInflections = inflections
	compile()
}

func SetSingular(inflections RegularSlice) {
	singularInflections = inflections
	compile()
}

func SetIrregular(inflections IrregularSlice) {
	irregularInflections = inflections
	compile()
}

func SetUncountable(inflections []string) {
	uncountableInflections = inflections
	compile()
}

func Plural(str string) string {
	for _, inflection := range compiledPluralMaps {
		if inflection.regexp.MatchString(str) {
			return inflection.regexp.ReplaceAllString(str, inflection.replace)
		}
	}
	return str
}

func Singular(str string) string {
	for _, inflection := range compiledSingularMaps {
		if inflection.regexp.MatchString(str) {
			return inflection.regexp.ReplaceAllString(str, inflection.replace)
		}
	}
	return str
}
