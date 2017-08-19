FROM scratch

ADD certs/ca-certificates.crt /etc/ssl/certs/
ADD go-crypto-average-service /
ADD aws/credentials /.aws/
ADD aws/config /.aws/

USER 1001

EXPOSE 8080

ENTRYPOINT ["./go-crypto-average-service"]