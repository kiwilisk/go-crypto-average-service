package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/kiwilisk/go-crypto-average-service/pb"
	"github.com/kiwilisk/go-crypto-average-service/s3bucket"
)

type FloatingAverageService interface {
	Load(symbol *string) (*floatingquotes.FloatingAverage, error)
}

type S3FloatingAverageService struct {
	bucket s3bucket.BinaryBucket
}

func NewS3FloatingAverageService(bucket s3bucket.BinaryBucket) *S3FloatingAverageService {
	return &S3FloatingAverageService{bucket}
}

func (service S3FloatingAverageService) Load(symbol *string) (*floatingquotes.FloatingAverage, error) {
	depotBytes, err := service.bucket.Download(symbol)
	if err != nil {
		return nil, err
	}
	floatingAverage := &floatingquotes.FloatingAverage{}
	err = proto.Unmarshal(depotBytes, floatingAverage)
	return floatingAverage, err
}
