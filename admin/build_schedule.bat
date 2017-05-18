SET GOOS=linux
SET GOARCH=amd64
go build -o sync-mysql-schedule ./main_schedule.go