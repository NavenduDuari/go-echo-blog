package main

import (
	"gitlab.com/yoursown/gocb/src/common/yoconst"
	"gitlab.com/yoursown/gocb/src/common/yolog"
	"gitlab.com/yoursown/gocb/src/yoblog/internal/store"
)

func main() {
	store.InitPostgressDB()

	db, err := store.GetPostgressDB()
	if err != nil {
		yolog.ErrorAndExit("GetPostgressDB failed", err)
	}

	err = InstallPsqlUUIDExtension(db)
	if err != nil {
		yolog.Error("Unable to install UUID extension on "+yoconst.PostgressYoblogDBName, err)
	}

	err = SetTimeZone(db)
	if err != nil {
		yolog.Error("Unable to set time zone on "+yoconst.PostgressYoblogDBName, err)
	}

	_, err = CreatePsqlTableYoblog(db)
	if err != nil {
		yolog.Error("Unable to create table: "+yoconst.PostgressYoblogTableBlog, err)
	}
}
