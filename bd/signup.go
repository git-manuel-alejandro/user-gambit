package bd

import (
	"fmt"

	"gihub.com/04-gambit/user-gambit/models"
	"gihub.com/04-gambit/user-gambit/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {

	fmt.Println("Comienza Registro")

	err := DBConect()

	if err != nil {
		return nil
	}

	defer Db.Close()

	// setencia := "INSERT INTO users (User_email , User__UUID , User_DateAdd)"

	setencia := fmt.Sprintf(`
    INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%v' , '%v', '%v')`,
		sig.UserEmail,
		sig.UserUUID,
		tools.FechaMySql())

	fmt.Println(setencia)

	_, err = Db.Exec(setencia)

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	fmt.Println("SignUp > Ejucion exitosa")

	return nil

}
