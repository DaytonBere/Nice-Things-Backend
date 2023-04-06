start:
	# Uses CompileDaemon to start the server
	CompileDaemon -command="./Nice-Things-Backend"

migrate-start:
	# Runs the migrate.go file 
	go run migrate/migrate.go
