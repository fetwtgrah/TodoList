package configs

type Config struct {
	Server struct {
		Port string
		Mode string
	}
	Database struct {
		Host     string
		Port     string
		Password string
		Dbname   string
		Sslmode  string
		User     string
	}
}
