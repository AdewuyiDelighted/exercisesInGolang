package main

func divisable() int {
	var answer = 0
	for count := 0; count <= 30; count++ {
		if count%3 == 0 {
			answer += count
		}

	}
	return answer
}
