package bd

import (
	"database/sql"
	"fmt"
	"os"

	"gihub.com/04-gambit/user-gambit/models"
	"gihub.com/04-gambit/user-gambit/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJSON
var err error
var Db *sql.DB

func ReadSecret() error {

	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DBConect() error {
	Db, err := sql.Open("mysql", ConsStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la base de datos")

	return nil

}

func ConsStr(claves models.SecretRDSJSON) string {

	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprint("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)

	fmt.Println(dsn)

	return dsn

}
