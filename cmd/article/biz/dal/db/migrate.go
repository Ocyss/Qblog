package db

func autoMigrate() error {
	return db.AutoMigrate(&Category{}, &Tag{}, &Article{})
}
