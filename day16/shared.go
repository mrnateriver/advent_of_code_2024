package day16

func justAWall(line string) bool {
	for i := range line {
		if line[i] != wall[0] {
			return false
		}
	}
	return true
}

const (
	empty = "."
	wall  = "#"
	start = "S"
	end   = "E"
)
