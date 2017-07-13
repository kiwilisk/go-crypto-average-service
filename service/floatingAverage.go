package service

import (
	"github.com/golang/protobuf/proto"
	"os"
	"github.com/kiwilisk/go-crypto-average-service/s3bucket"
	"github.com/kiwilisk/go-crypto-average-service/pb"
)

func LoadFloatingAverage(symbol *string) (*floatingquotes.FloatingAverage, error) {
	bucketName := os.Getenv("s3.bucketName")
	depotBytes, err := s3bucket.Download(&bucketName, symbol)
	if err != nil {
		return nil, err
	}
	floatingAverage := &floatingquotes.FloatingAverage{}
	err = proto.Unmarshal(depotBytes, floatingAverage)
	return floatingAverage, err
}
