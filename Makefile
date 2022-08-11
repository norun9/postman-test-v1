linux-setup:
	go install github.com/volatiletech/sqlboiler/v4@v4.11.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.11.0
	go install github.com/pressly/goose/cmd/goose@latest
	go install github.com/google/wire/cmd/wire@latest
