# build binary
FROM golang:1.13.5 AS build

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/simple-redirect-service
ADD . .

RUN dep ensure

RUN ./testall.sh

RUN GOOS=linux CGO_ENABLED=0 go build

# create runnable image
FROM scratch

COPY --from=build /go/src/simple-redirect-service/simple-redirect-service /simple-redirect-service

ENTRYPOINT [ "/simple-redirect-service" ]
