package service

import (
	"api-produto/entity"
	"log"
)

func (ps *produto_service) GetUser(login *string) *entity.User {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT login, senha FROM usuario WHERE login = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	user := entity.User{}

	err = stmt.QueryRow(login).Scan(&user.Username, &user.Senha)
	if err != nil {
		log.Println("error: cannot find produto", err.Error())
	}

	return &user
}
