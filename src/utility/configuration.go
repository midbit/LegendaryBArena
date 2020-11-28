package utility

import "os"

func GetToken() string {
	return os.Getenv("TOKEN")
}

func GetMySQLHost() string {
	var host string
	host = os.Getenv("MYSQL_HOST")
	if host == "" {
		return "127.0.0.1"
	}
	return host
}

func GetMySQLPort() string {
	var port string
	port = os.Getenv("MYSQL_PORT")
	if port == "" {
		return "3306"
	}
	return port
}

func GetMySQLUser() string {
	var user string
	user = os.Getenv("MYSQL_USER")
	if user == "" {
		return "root"
	}
	return user
}

func GetMySQLPassword() string {
	var password string
	password = os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		return "root"
	}
	return password
}

func GetMySQLDatabase() string {
	var database string
	database = os.Getenv("MYSQL_DATABASE")
	if database == "" {
		return "LegendaryABeam"
	}
	return database
}
