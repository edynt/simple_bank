<!-- migration db -->
1. migrate create -ext sql -dir db/migration -seq init_schema

2. make migrateup

3. go mod init github.com/edynnt/simple_bank