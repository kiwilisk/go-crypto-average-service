package s3bucket

import (
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Download(bucketName *string, symbol *string) ([]byte, error) {
	key := createS3KeyFor(*symbol)
	sess := createSession()
	downloader := s3manager.NewDownloader(sess)
	buffer := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: bucketName,
			Key:    key,
		})
	if err != nil {
		err = fmt.Errorf("Unable to download item with symbol %s from bucket %s, %v",
			aws.StringValue(symbol), aws.StringValue(bucketName), err)
	}
	return buffer.Bytes(), err
}

func createSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func createS3KeyFor(symbol string) *string {
	bytes := []byte(symbol)
	hexEncodedSymbol := hex.EncodeToString(bytes)
	key := hexEncodedSymbol + "_" + symbol
	return &key
}
