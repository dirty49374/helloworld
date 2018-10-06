FROM golang as builder

WORKDIR /

COPY main.go /
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o helloworld .


#####################################
FROM scratch

WORKDIR /
COPY --from=builder /helloworld /

EXPOSE 80
ENTRYPOINT [ "/helloworld" ]

