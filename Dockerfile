FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . ./

RUN go build -o /app/omesrv


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/omesrv ./
COPY --from=builder /app/cfg/omesrv.yaml ./

ARG OME_REDIS_STORE_POOL
ARG OME_REDIS_STORE_ADDR
ARG OME_REDIS_STORE_PASS

ENV OME_REDIS_STORE_POOL=$OME_REDIS_STORE_POOL
ENV OME_REDIS_STORE_ADDR=$OME_REDIS_STORE_ADDR
ENV OME_REDIS_STORE_PASS=$OME_REDIS_STORE_PASS

EXPOSE 9080

CMD [ "/app/omesrv" ]