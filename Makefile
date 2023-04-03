start:
	# Uses CompileDaemon to start the server
	CompileDaemon -command="./Nice-Things-Backend"

migrate-up:
	# Runs the migrate.go file to auto migrate all the tables
	go run migrate/migrate.go