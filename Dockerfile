FROM golang:1.15 AS build
WORKDIR /go/src/molde
COPY . .
SHELL ["/bin/bash", "-c"]
RUN go get -d -v ./...
RUN go build -o ./dist/molde -v ./cmd/molde/

FROM node:14.15.4
WORKDIR /app
COPY --from=build /go/src/molde/dist/ .
COPY --from=build /go/src/molde/package.json package.json
RUN npm install
