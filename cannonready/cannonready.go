package cannonready

type Gunner struct {
	Name  string
	Reply string
}

func CannonReady(gunners []Gunner) string {
	for _, gunner := range gunners {
		if gunner.Reply != "aye" {
			return "Shiver me timbers!"
		}
	}
	return "Fire!"
}
