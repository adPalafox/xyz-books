package constant

const (
	ResponseOKMessage               = "OK"
	ResponseInvalidParameterMessage = "Invalid Parameter"
	ResponseInvalidArgumentMessage  = "Invalid Argument"
	ResponseNotfoundMessage         = "Not found"
	ResponseInternalServerMessage   = "Internal Server Error"
)

const (
	ProgramDatabaseMigrationDirectory = "./sqlite.db"
	ProgramDatabaseAccessDirectory    = "../../sqlite.db"
	ProgramSeedBookDirectory          = "cmd/db/seeds/books_seed_data.json"
	ProgramSeedPublisherDirectory     = "cmd/db/seeds/publishers_seed_data.json"
)

const (
	LogStartMessage             = "Start"
	LogFinishMessage            = "Finish"
	LogDatabaseSuccessMigration = "Successfully migrated the database"
	LogDatabaseRunningMigration = "Running migration"
	LogDatabaseFailureToConnect = "Failed to connect to the database"
	LogDatabaseFailureToMigrate = "Failed to migrate database"
	LogDatabaseFailureToSeed    = "Failed to seed database"
)

const (
	SortByTitle           = "title"
	SortByListPrice       = "list_price"
	SortByPublicationYear = "publication_year"
	OrderByAsc            = "asc"
	OrderByDsc            = "desc"
)
