package s3bucket

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BinaryBucket interface {
	Download(key *string) ([]byte, error)
}

type S3BinaryBucket struct {
	downloader  *s3manager.Downloader
	keyProvider S3KeyProvider
	bucketName  *string
}

func NewS3BinaryBucket(downloader *s3manager.Downloader, keyProvider S3KeyProvider, bucketName *string) *S3BinaryBucket {
	return &S3BinaryBucket{downloader, keyProvider, bucketName}
}

func (binaryBucket *S3BinaryBucket) Download(key *string) ([]byte, error) {
	downloader := binaryBucket.downloader
	keyProvider := binaryBucket.keyProvider
	buffer := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: binaryBucket.bucketName,
			Key:    keyProvider.CreateKeyFor(*key),
		})
	if err != nil {
		err = fmt.Errorf("Unable to download item with symbol %s from bucket %s, %v",
			aws.StringValue(key), aws.StringValue(binaryBucket.bucketName), err)
	}
	return buffer.Bytes(), err
}
