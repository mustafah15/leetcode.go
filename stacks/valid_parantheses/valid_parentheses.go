package stacks

func isValid(s string) bool {
	st := stack{
		store: []rune{},
	}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			st.store = append(st.store, ch)
		} else {
			if ok := st.remove(ch); !ok {
				return false
			}
		}
	}

	return len(st.store) == 0
}

type stack struct {
	store []rune
}

func (s *stack) remove(ch rune) bool {
	if len(s.store) == 0 {
		return false
	}
	switch ch {
	case ')':
		if s.store[len(s.store)-1] != '(' {
			return false
		}

	case '}':
		if s.store[len(s.store)-1] != '{' {
			return false
		}

	case ']':
		if s.store[len(s.store)-1] != '[' {
			return false
		}
	}
	s.store = s.store[:len(s.store)-1]
	return true
}
