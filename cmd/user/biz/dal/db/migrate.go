package db

func autoMigrate() error {
	return db.AutoMigrate(&User{})
}
