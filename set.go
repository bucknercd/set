package set

type Set struct {
	Set map[string]bool
}

func NewSet() *Set {
	s := Set{}
	s.Set = make(map[string]bool)
	return &s
}


func (s *Set) Add(element string) {
	s.Set[element] = true
}

func (s *Set) Delete(element string) {
	delete(s.Set, element)
}

func (s *Set) Size() int {
	return len(s.Set)
}

func (s *Set) Exists(element string) bool {
	return s.Set[element]
}