package config

type Config struct {
	DB_url   string `json:"db_url"`
	Username string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"
