package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, el := range l {
		set[el] = true
	}
	return set
}

func (s Set) String() string {
	retStr := ""

	if len(s) == 0 {
		return "{}"
	}
	for el := range s {
		retStr += `"` + el + `", `
	}
	return "{" + retStr[:len(retStr)-2] + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for el := range s1 {
		if !s2.Has(el) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for el := range s1 {
		if s2.Has(el) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && len(s1) == len(s2)
}

func Intersection(s1, s2 Set) Set {
	set := New()
	for el := range s1 {
		if s2.Has(el) {
			set.Add(el)
		}
	}
	return set
}

func Difference(s1, s2 Set) Set {
	set := New()
	for el := range s1 {
		if !s2.Has(el) {
			set.Add(el)
		}
	}
	return set
}

func Union(s1, s2 Set) Set {
	set := New()
	for _, s := range []Set{s1, s2} {
		for el := range s {
			set.Add(el)
		}
	}
	return set
}
