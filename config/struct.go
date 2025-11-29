package config

type Config struct {
	Postgres Postgres `json:"postgres"`
}


type Postgres struct {
	Host     string `json:"host" default:"localhost"`
	Port     string `json:"port" default:"54321"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	DB       string `json:"db" default:"noroask"`
	Timezone string `json:"time_zone" default:"UTC"`
	SSLMode  string `json:"ssl_mode" default:"disable"`
}