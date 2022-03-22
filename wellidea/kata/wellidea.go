package kata

func Well(ss []string) string {
	goods := 0
	for _, s := range ss {
		if s == "good" {
			goods += 1
		}
	}
	switch goods {
	case 0:
		return "Fail!"
	case 1:
		return "Publish!"
	case 2:
		return "Publish!"
	default:
		return "I smell a series!"
	}
}
