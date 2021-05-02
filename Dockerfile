FROM golang:1.13-alpine3.11 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o employee-list .

FROM alpine:3.11.3
COPY --from=builder /app/employee-list .
ENV EMPLOYEE_STORE "serviceb"
EXPOSE 8081
CMD [ "./employee-list" ]
