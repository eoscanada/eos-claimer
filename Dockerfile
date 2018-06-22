FROM gcr.io/cloud-builders/go as builder

ENV CGO_ENABLED 0
ENV PKG /root/go/src/github.com/eoscanada/eos-claimer
RUN mkdir -p $PKG
COPY . $PKG
RUN cd $PKG \
    && go get -v -t . \
    && go test -v \
    && go build -v -o /eos-claimer


FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /eos-claimer /app/eos-claimer

ENV EOS_CLAIMER_URL https://mainnet.eoscanada.com
ENV EOS_CLAIMER_OWNER eoscanadacom
ENV EOS_CLAIMER_PERMISSION claimer
ENV EOS_CLAIMER_KEY none

CMD /app/eos-claimer
