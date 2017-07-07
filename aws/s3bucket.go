package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func PrintBuckets() {
	sess := createSession()
	s3client := s3.New(sess)
	result := retrieveListBucketResult(s3client)
	printBucketsFrom(result)
}

func createSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func retrieveListBucketResult(s3client *s3.S3) *s3.ListBucketsOutput {
	result, err := s3client.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}
	return result
}

func printBucketsFrom(result *s3.ListBucketsOutput) {
	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
