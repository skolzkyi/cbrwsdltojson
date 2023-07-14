FROM golang:1.20.5 as build

ENV BIN_FILE /opt/cbrwsdltojson/integrationtests
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

#RUN mkdir /opt/cbrwsdltojson

ARG LDFLAGS
RUN CGO_ENABLED=0  go test -c \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} integrationtests/*

CMD ${BIN_FILE}

FROM alpine:3.9

LABEL ORGANIZATION=""
LABEL SERVICE="integration_tests"
LABEL MAINTAINERS="skolzkyi@gmail.com"

ENV BIN_FILE "/opt/cbrwsdltojson/integrationtests"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

RUN chmod 555 ${BIN_FILE}
#RUN chmod ugo+x ${BIN_FILE}

ENV CONFIG_FILE /etc/cbrwsdltojson/config.env
COPY ./configs/config.env ${CONFIG_FILE}

RUN mkdir /deployLogs

CMD  ${BIN_FILE} -config ${CONFIG_FILE} -test.v > /deployLogs/integrationTestsCBRWSToJSONLog.log