package fortune

import "os/exec"

// GetFortune returns a randomly generated fortune using BSD `fortune`
func GetFortune(offensive bool) (string, error) {
	args := []string{}
	if offensive {
		args = []string{"-o"}
	}

	cmd := exec.Command("fortune", args...)
	outBytes, err := cmd.Output()

	return string(outBytes), err
}
