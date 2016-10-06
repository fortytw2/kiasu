package hubris

// A Config is a tuple of unique values attached to a scrape
type Config struct {
	Country string
	Type    string
	Scrape  string
}

// AnyCountry is a config in which country doesn't matter
var AnyCountry = "any"
