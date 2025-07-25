MIGRATIONS_FOLDER=db/migrations
DB_URL=root:auth-dev-pw@tcp(localhost:3306)/auth_dev

# Creates new migration file with the current timestamp
migration-create: # gmake migration-create name="create-table-user"
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

# Migrate the DB to the most recent version available
migrate-up:	# gmake migrate-up
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

# Roll back the version by 1
migrate-down:	# gmake migrate-down
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

# Roll back all migrations and reset database 
migrate-reset:	# gmake migrate-reset
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" reset

# Show current migration status
migrate-status:	# gmake migrate-status
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) 
migrate-redo: # gmake migrate-redo
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" redo

# Run specific migration version
migrate-to:	# gmake migrate-to version=1234567890
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version
migrate-down-to:	# gmake migrate-down-to version=1234567890
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down to $(version)

# Force a specific migration version 
migrate-force:	# gmake migrate-force version=1234567890
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" force $(version)

# Print Goose help
migrate-help:	# gmake migrate-help
	goose -h