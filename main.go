package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/kiwilisk/go-crypto-average-service/s3bucket"
	"github.com/kiwilisk/go-crypto-average-service/service"
	"log"
	"net/http"
	"os"
)

const envS3BucketName = "s3.bucketName"

func main() {
	binaryBucket := createS3Bucket()
	floatingAverageService := service.NewS3FloatingAverageService(binaryBucket)
	currencyHandler := service.NewCurrencyHandler(floatingAverageService)
	router := createRouter(currencyHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createS3Bucket() *s3bucket.S3BinaryBucket {
	sharedSession := createSession()
	downloader := s3manager.NewDownloader(sharedSession)
	keyProvider := s3bucket.NewHexKeyProvider()
	bucketName := os.Getenv(envS3BucketName)
	binaryBucket := s3bucket.NewS3BinaryBucket(downloader, keyProvider, &bucketName)
	return binaryBucket
}

func createRouter(currencyHandler *service.CurrencyHandler) *mux.Router {
	routes := service.Routes{
		service.Route{
			Name:        "CurrencyAverageQuotes",
			Method:      "GET",
			Pattern:     "/currencies/{symbol}",
			HandlerFunc: currencyHandler.Handle,
		}}
	return service.NewRouter(routes)
}

func createSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}
