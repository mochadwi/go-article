package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/mochadwi/go-article/middleware"
	"github.com/mochadwi/go-article/models"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	articleDeliverEcho "github.com/mochadwi/go-article/article/delivery/http_echo"
	articleRepo "github.com/mochadwi/go-article/article/repository"
	articleUcase "github.com/mochadwi/go-article/article/usecase"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	//fmt.Println(connection)
	dbConn, err := gorm.Open(`postgres`, connection)
	dbConn.LogMode(true)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = dbConn.DB().Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Migrate the schema
	dbConn.AutoMigrate(&models.Article{})

	middL := middleware.InitMiddleware()

	e := echo.New()
	e.Use(middL.CORS)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	ar := articleRepo.NewGormsqlArticleRepository(dbConn)
	au := articleUcase.NewArticleUsecase(ar, timeoutContext)

	articleDeliverEcho.NewArticleHttpEchoHandler(e, au)

	e.Start(viper.GetString("server.address"))
}
