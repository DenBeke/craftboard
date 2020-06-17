FROM node:latest AS front

COPY /app /app
WORKDIR /app


RUN yarn install
RUN yarn build


FROM golang:latest AS back

WORKDIR /

COPY / /random_work_dir

WORKDIR /random_work_dir

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /craftboard ./cmd/craftboard


FROM alpine:latest

WORKDIR /

COPY --from=front /app/dist /craftboard/app/dist
COPY --from=back /craftboard /craftboard/
COPY board_sample.json /craftboard/board_sample.json

WORKDIR /craftboard

RUN chmod +x /craftboard

EXPOSE 1234
VOLUME ["/craftboard"]

CMD ["./craftboard"]