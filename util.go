package GoBot

func joinMessage(strs ...string) string {
	var message string
	for _, str := range strs {
		message += str
	}
	return message
}

func containsInSlices(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
