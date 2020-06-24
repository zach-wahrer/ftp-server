module github.com/zachtheclimber/ftp-server

go 1.13

require internal/server v1.0.0

replace internal/server => ./internal/server

require internal/ui v1.0.0

replace internal/ui => ./internal/ui
