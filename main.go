package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	articleDeliverHttpEcho "github.com/mochadwi/go-article/features/article/delivery/httpecho"
	articleRepo "github.com/mochadwi/go-article/features/article/repository"
	articleUcase "github.com/mochadwi/go-article/features/article/usecase"
	ratingDeliverHttpEcho "github.com/mochadwi/go-article/features/rating/delivery/httpecho"
	ratingRepo "github.com/mochadwi/go-article/features/rating/repository"
	ratingUcase "github.com/mochadwi/go-article/features/rating/usecase"
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

	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

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
		err = dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Migrate the schema
	dbConn.AutoMigrate(&models.Article{})
	dbConn.AutoMigrate(&models.Rating{})

	middL := middleware.InitMiddleware()

	echoStart := echo.New()
	echoStart.Use(middL.CORS)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	articleR := articleRepo.NewGormsqlArticleRepository(dbConn)
	articleU := articleUcase.NewArticleUsecase(articleR, timeoutContext)
	articleDeliverHttpEcho.NewArticleHttpEchoHandler(echoStart, articleU)

	ratingR := ratingRepo.NewGormsqlRatingRepository(dbConn)
	ratingU := ratingUcase.NewRatingUsecase(ratingR, timeoutContext)
	ratingDeliverHttpEcho.NewRatingHttpEchoHandler(echoStart, ratingU)

	echoStart.Start(viper.GetString("server.address"))
}
