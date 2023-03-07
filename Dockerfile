# golang base image that's so minimum
FROM golang:1.16-alpine

# set the working directory
WORKDIR .

# copy (source code) to (working directory in container)
COPY . .

# build the go app
RUN go build -o main .

# expose port 8080
EXPOSE 8080

# run the go app
CMD ["./main"]