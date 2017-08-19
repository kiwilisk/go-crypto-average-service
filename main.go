package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/kiwilisk/go-crypto-average-service/s3bucket"
	"github.com/kiwilisk/go-crypto-average-service/service"
	"log"
	"net/http"
	"flag"
)

const bucketFlagName = "bucket"

func main() {
	log.Println("Starting go-crypto-average-service, listening on 8080")
	bucketName := flag.String(bucketFlagName,  "someBucket","the s3 bucket name to retrieve data")
	flag.Parse()

	binaryBucket := createS3Bucket(bucketName)
	floatingAverageService := service.NewS3FloatingAverageService(binaryBucket)
	currencyHandler := service.NewCurrencyHandler(floatingAverageService)
	router := createRouter(currencyHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createS3Bucket(bucketName *string) *s3bucket.S3BinaryBucket {
	sharedSession := createSession()
	downloader := s3manager.NewDownloader(sharedSession)
	keyProvider := s3bucket.NewHexKeyProvider()
	binaryBucket := s3bucket.NewS3BinaryBucket(downloader, keyProvider, bucketName)
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
