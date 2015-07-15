Inflection
=========

Inflection pluralizes and singularizes English nouns

## Basic Usage

```
inflection.Plural("person") => people
inflection.Plural("Person") => People
inflection.Plural("PERSON") => PEOPLE

inflection.Singularize("people") => "person"
inflection.Singularize("People") => "Person"
inflection.Singularize("PEOPLE") => "PERSON"

inflection.Plural("FancyPerson") => FancydPeople
inflection.Singularize("FancyPeople") => FancydPerson
```


## Register Rules

```
inflection.AddUncountable("fish")
inflection.AddIrregular("person", "people")
inflection.AddPlural("(bu)s$", "${1}ses") // bus => buses / BUS => BUSES / Bus => Buses
inflection.AddSingular("(bus)(es)?$", "${1}") // buses => bus / Buses => Bus / BUSES => BUS
```

