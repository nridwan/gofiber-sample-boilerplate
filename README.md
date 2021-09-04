
  
# Go Fiber Boilerplate    
## Installation
 1. Install golang 1.17
 2. open project
 3. run in terminal/bash/powershell/anything that works in that OS
     `go get` 
     `go mod tidy`
 4. proceed install SQLBoiler by using:
    `go install github.com/volatiletech/sqlboiler/v4@latest`
    `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest`
    *note: if using postgres use sqlboiler-psql instead of sqlboiler-mysql. though i'm not sure if the migration works in postgres or need adjustment, just try it.
 5. copy `sqlboiler.toml.example` to `sqlboiler.toml` 
 6. copy `.env.example` to `.env`
 7. edit database configuration in both `sqlboiler.toml` and `.env`
 8. run migration by running `go run .\migration\migrate.go --up`
 9. install Fiber CLI for hot reload support 
    `go get -u github.com/gofiber/cli/fiber`
 10. run the server by `fiber dev`
 11. any changes to code will be automatically recompiled
## How Create Migrations
1. run the command `./make_migration.sh {name}` or `.\make_migration.ps1 {name}`, then `timestamp_{name}.up.sql` and `timestamp_{name}.down.sql` is created
3. in up.sql insert any sql query to add/modify tables
4. in down.sql insert any sql query that revert anything did in up.sql
## How to Use Migrate.go
- Apply migration: `go run .\migration\migrate.go --up`
  this will apply all available migrations
- Revert migration: `go run .\migration\migrate.go --down --step={number}`
  replace {number} with how many version needed to revert
  if not specified, only revert 1 version
- Force migration version: `go run .\migration\migrate.go --force={version_number}`
- Check current version: `go run .\migration\migrate.go`
## Build Production Binary
- Build for current OS: `.\build.ps1` | `./build.sh`
- Build for other OS: `.\build.ps1 {OS} {ARCH}` | `./build.sh {OS} {ARCH}`
- replace {OS} and {ARCH} with valid value
  [See OS and ARCH list here](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
  example, windows creating linux x64 binary using `.\build.ps1 linux amd64`
## Notes
 - there's only checking mechanism for this boilerplate, proceed by entering data manually to database before testing.
 - password used for `users` table is using `bcrypt`, use any online generator or just get some data from laravel projects if any
## Coming Soon / Todos (not sure when, don't ask)
- App Login
- Cleanup Junk JWT Token
## Disclaimer
- All these tutorials are made for windows and powershell, but there's bash script just in case for those who's using *nix based machine.