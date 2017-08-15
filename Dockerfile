FROM golang:alpine
EXPOSE 8080

ENV USER_NAME crypto_user
ENV APP_HOME /home/$USER_NAME/app
ENV s3.bucketName floating-average-crypto

RUN adduser -D -u 1000 $USER_NAME
RUN mkdir $APP_HOME

ADD go-crypto-average-service $APP_HOME
ADD aws/credentials /home/$USER_NAME/.aws/
ADD aws/config /home/$USER_NAME/.aws/

RUN chown $USER_NAME $APP_HOME/go-crypto-average-service

USER $USER_NAME
WORKDIR $APP_HOME

ENTRYPOINT ["./go-crypto-average-service"]