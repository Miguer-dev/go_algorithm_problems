package stack

func isValid(s string) bool {
	stack := LIFO{}
	openClose := make(map[rune]rune)
	openClose[')'] = '('
	openClose[']'] = '['
	openClose['}'] = '{'

	for _, value := range s {
		if _, exits := openClose[value]; !exits {
			//if ( o [ o {
			stack.push(value)
		} else if len(stack) != 0 {
			//if ) o ] o }
			open := stack.pop()
			if open != openClose[value] {
				return false
			}
		} else {
			return false
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
