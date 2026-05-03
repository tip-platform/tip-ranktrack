package error

type InvalidPlayerNameError struct{}

func (e InvalidPlayerNameError) Error() string {
	return "player name cannot be empty"
}
