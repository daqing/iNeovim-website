FROM reg.mindwork.cool/golang:1.27-alpine AS builder
WORKDIR /app
COPY . /app
RUN go build -o ./bin/app .

FROM reg.mindwork.cool/alpine
WORKDIR /app
COPY --from=builder /app/bin/app /app
COPY --from=builder /app/db /app/db
COPY --from=builder /app/static /app/static

ENV AIRWAY_ENV=production
ENV LISTEN=:1905
ENV STATIC_PATH=/app/static
ENV TZ="Asia/Shanghai"

EXPOSE 1905

CMD ["/app/app"]
