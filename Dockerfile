FROM golang as builder

# Add Maintainer Info
LABEL maintainer="Muhammad Indrawan <me@indra.codes>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/aboutindra/fokusngoding-compiler-javascript

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/github.com/aboutindra/fokusngoding-compiler-javascript /app/fokusngoding-compiler-javascript
WORKDIR "/app/fokusngoding-compiler-javascript"
EXPOSE 3500
ENTRYPOINT ./fc-javascript
