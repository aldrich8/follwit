package config

import "fmt"

func DBConfiguration(env string) string {
	dbHost := fmt.Sprintf("%s_postgres_db", env)
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", dbHost, "postgres", "postgres", "follwit", "5432", "disable")
}
