package mockgen

func SaveUser(db Database, key, value string) error {
	return db.Save(key, value)
}

func GetUser(db Database, key string) (string, error) {
	return db.Load(key)
}
