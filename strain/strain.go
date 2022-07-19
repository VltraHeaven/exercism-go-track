package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (newInts Ints) {
	for n := range i {
        if filter(i[n]) {
            newInts = append(newInts, i[n])
        }
    }
	return
}

func (i Ints) Discard(filter func(int) bool) (newInts Ints) {
	for n := range i {
        if !filter(i[n]) {
            newInts = append(newInts, i[n])
        }
    }
	return
}

func (l Lists) Keep(filter func([]int) bool) (newLists Lists) {
	for n := range l {
        if filter(l[n]) {
            newLists = append(newLists, l[n])
        }
    }
	return
}

func (s Strings) Keep(filter func(string) bool) (newStrings Strings) {
	for n := range s {
        if filter(s[n]) {
            newStrings = append(newStrings, s[n])
        }
    }
	return
}
