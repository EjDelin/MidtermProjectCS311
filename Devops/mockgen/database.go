package mockgen

//go:generate mockgen -source=database.go -destination=mockgen_database.go -package=mockgen

type Database interface {
	Save(key string, value string) error
	Load(key string) (string, error)
}
