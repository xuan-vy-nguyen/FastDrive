FROM golang:1.10.4

RUN go get github.com/xuan-vy-nguyen/SE_Project01

EXPOSE 8080

CMD ["SE_Project01", "run"]