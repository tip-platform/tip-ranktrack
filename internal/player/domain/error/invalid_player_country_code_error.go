package error

type InvalidPlayerCountryCodeError struct{}

func (e InvalidPlayerCountryCodeError) Error() string {
	return "player country code cannot be empty"
}
