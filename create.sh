# STEP 1: Create Files And Directories
mkdir -p "Configs"
mkdir -p "Models"
mkdir -p "Utils/Errors"
mkdir -p "Services"
mkdir -p "App"
mkdir -p "Controllers/Ping"

touch Configs/setup.env
touch Configs/db.go
touch README.md
touch Utils/Errors/rest_errors.go
touch Controllers/Ping/ping_controller.go
touch .gitignore

# STEP 2: Get All Of Variables For ENV ...
read -p 'Project Name: '  PROJECT_NAME
read -p 'Server Port: '  PORT
read -p 'Database User: '  DB_USER
read -p 'Database Password: '  DB_PASSWORD
read -p 'Database Name: '  DB_NAME
read -p 'Datebase Port: '  DB_PORT
read -p 'Database Host: ' DB_HOST

# STEP 3: Fill Up Files With Data
cat <<EOT >> Configs/setup.env
PORT= $PORT
DB_USER= $DB_USER
DB_PASSWORD= $DB_PASSWORD
DB_NAME= $DB_NAME
DB_HOST= $DB_HOST
EOT

cat <<EOT >> main.go
package main
import(
    "github.com/kasrashrz/Golang/App"
)
func main(){
    app.StartApp()
}
EOT

cat <<EOT >> App/application.go
package app
import "github.com/gin-gonic/gin"

var (
    router = gin.Default()
)

func StartApp(){
    mapURLs()
    router.Run(":$PORT")
}

EOT

cat <<EOT >> App/url_mapping.go
package app
import 
(
    "github.com/kasrashrz/Golang/Controllers"
)
func mapURLs (){
    router.GET("/ping", ping.Ping)
}
EOT

cat <<EOT >> Controllers/ping_controller.go
package ping
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func Ping(ctx *gin.Context){
    ctx.String(http.StatusOK, "Pong\n")
}
EOT

cat <<EOT >> README.md
# $PROJECT_NAME project started ;D
EOT

go mod init "github.com/kasrashrz/Golang"
go run main.go