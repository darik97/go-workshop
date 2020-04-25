package api

// Client interface with 3-rd party joke API
type Client interface {
	// GetJoke return one joke
	GetJoke() (*JokeResponse, error)
}
