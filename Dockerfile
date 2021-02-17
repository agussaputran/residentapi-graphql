FROM golang:1.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o resident-graphql .
CMD ["/app/resident-graphql"]