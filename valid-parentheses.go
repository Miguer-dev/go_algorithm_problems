package main

func isValid(s string) bool {
	queue := LIFO{}
	openClose := make(map[rune]rune)
	openClose[')'] = '('
	openClose[']'] = '['
	openClose['}'] = '{'

	for _, value := range s {
		if _, exits := openClose[value]; !exits {
			//if ( o [ o {
			queue.push(value)
		} else if len(queue.stack) != 0 {
			//if ) o ] o }
			open := queue.pop()
			if open != openClose[value] {
				return false
			}
		} else {
			return false
		}
	}

	if len(queue.stack) != 0 {
		return false
	}

	return true
}
