package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	httpDeliver "github.com/mochadwi/go-article/article/delivery/http"
	articleRepo "github.com/mochadwi/go-article/article/repository"
	articleUcase "github.com/mochadwi/go-article/article/usecase"
	"github.com/mochadwi/go-article/middleware"
	"github.com/mochadwi/go-article/models"
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
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.DB().Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer dbConn.Close()

	// Migrate the schema
	dbConn.AutoMigrate(&models.Article{})

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	//ar := articleRepo.NewMysqlArticleRepository(dbConn)
	ar := articleRepo.NewGormsqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := articleUcase.NewArticleUsecase(ar, timeoutContext)
	httpDeliver.NewArticleHttpHandler(e, au)

	e.Start(viper.GetString("server.address"))
}
