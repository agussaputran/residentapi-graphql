package helper

import (
	"os"
)

// Token global
var Token string

// ReqBody global
var ReqBody string

// SetEnvVar func
func SetEnvVar() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "r3m3mb3r34")
	os.Setenv("DB_NAME", "resident-graphql")
	os.Setenv("DB_PORT", "5000")
	os.Setenv("DB_SSL", "disable")
	os.Setenv("DB_TIMEZONE", "Asia/Shanghai")
	os.Setenv("JWT_SECRET", "kodebom")
	os.Setenv("RAJAONGKIR_APIKEY", "366cf34aabf54556c13fd48951aadf3a")
	os.Setenv("DSN", "https://a01e543958d44f51b3860e928ca45980@o515018.ingest.sentry.io/5619088")
}
