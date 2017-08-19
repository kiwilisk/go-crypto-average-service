# go-crypto-average-service
Retrieve moving average information of a crypto currency.

This project serves the purpose of learning go.
Can only be used in conjunction with [go-crypto-average-service](https://github.com/kiwilisk/go-crypto-average-service), which stores the data that should be fetched with this service.

This simple rest service will retrieve a [protobuf](https://developers.google.com/protocol-buffers/) file from [S3]((https://aws.amazon.com/s3/)) for the given crypto currency (i.e. bitcoin) and return a JSON response, containing the latest moving average and the quotes for the last week.

Dependency management with [dep](https://github.com/golang/dep).

Usage with docker (just a reminder to me):
* copy credentials file from ~/.aws directory to this aws folder
* execute build.sh
* execute docker run -p 8080:8080 go-crypto-average-service -bucket=yourBucketName