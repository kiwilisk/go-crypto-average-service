package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-crypto-average-service/myaws"
	"go-crypto-average-service/pb"
	"os"
)

func main() {
	myaws.PrintBuckets()
	bucketName, key := bucketNameAndKeyFromArgs()

	depotBytes := myaws.Download(&bucketName, &key)
	depot := &floatingquotes.Depot{}
	err := proto.Unmarshal(depotBytes, depot)

	if err != nil {
		exitErrorf("Failed to download", err)
	}
	averages := depot.GetFloatingAverages()
	fmt.Println("Downloaded proto", averages[0])
}
func bucketNameAndKeyFromArgs() (string, string) {
	arguments := os.Args[1:]
	bucketName, key := arguments[0], arguments[1]
	return bucketName, key
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
