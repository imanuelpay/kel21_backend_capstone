package user

import (
	"backend_capstone/services/user"
	"backend_capstone/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) user.Repository {
	var transactionsRepo user.Repository

	if dbCon.Driver == utils.Postgres {
		transactionsRepo = NewPostgresRepository(dbCon.Postgres)
	}
	if dbCon.Driver == utils.MySQL {
		transactionsRepo = NewPostgresRepository(dbCon.MySQL)
	}

	return transactionsRepo
}
