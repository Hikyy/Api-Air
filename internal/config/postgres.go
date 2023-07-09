package config

//"github.com/joho/godotenv"

func Postgres() map[string]string {

	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Println("Failed to load .env file:", err)
	// }

	// var (
	// 	host     = os.Getenv("POSTGRES_HOST")
	// 	port     = os.Getenv("POSTGRES_PORT")
	// 	dbuser   = os.Getenv("POSTGRES_USER")
	// 	password = os.Getenv("POSTGRES_PASSWORD")
	// 	dbname   = os.Getenv("POSTGRES_DB")
	// )

	var (
		host     = "localhost"
		port     = "3307"
		dbuser   = "root"
		password = "root"
		dbname   = "postgres"
	)

	return map[string]string{
		"host":     host,
		"port":     port,
		"dbuser":   dbuser,
		"password": password,
		"dbname":   dbname,
	}
}
