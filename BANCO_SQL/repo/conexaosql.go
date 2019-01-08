package repo

import (
	/*
		github.com/go-sql-driver/mysql não é diretamente usado pela aplicacao
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db é um ponteiro para o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL abre a conexao com DB SQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil

	Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}

	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
