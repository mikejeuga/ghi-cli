package models

type Config struct {
	Repo     string
	Username string
	Tk       string
	HAccept  string
}

func NewConfig() *Config {
	repo := "ghi-cli/issues"
	username := ""

	return &Config{
		Repo:     repo,
		Username: username,
		Tk:       "",
		HAccept:  "application/vnd.github.v3+json",
	}
}
