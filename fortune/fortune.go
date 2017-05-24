package fortune

import "os/exec"

// Length is an alias for int, for use as an enum
type Length int

// Fortune length options
const (
	Short Length = iota
	Long
	All
)

// GetFortune returns a randomly generated fortune using BSD `fortune`
func GetFortune(offensive bool, length Length, passthroughOptions []string) (string, error) {
	args := []string{}
	if offensive {
		args = append(args, "-o")
	}

	switch length {
	case Short:
		args = append(args, "-s")
	case Long:
		args = append(args, "-l")
	case All:
		// noop
	}

	args = append(args, passthroughOptions...)

	cmd := exec.Command("fortune", args...)
	outBytes, err := cmd.Output()

	return string(outBytes), err
}
