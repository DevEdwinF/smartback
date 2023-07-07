package config

import "database/sql"

func GetDBConnection() (*sql.DB, error) {
	connString := "server=34.148.172.87;user id=sa;password=*T3cnol0gi+*;database=kactus"

	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
