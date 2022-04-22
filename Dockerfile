FROM golang:1.18

RUN mkdir /usr/bin/rest-api-template

ADD go.mod go.sum Makefile /usr/bin/rest-api-template
ADD internal/ /usr/bin/rest-api-template
ADD vendor/ /usr/bin/rest-api-template
ADD cmd/ /usr/bin/rest-api-template

WORKDIR /usr/bin/rest-api-template

CMD ["make", "build"]
CMD ["make", "run"]
