FROM golang:1.13.5

RUN go get github.com/xuan-vy-nguyen/SE_Project01

EXPOSE 8080

CMD ["SE_Project01", "run"]