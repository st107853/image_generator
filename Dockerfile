FROM golang:1.23.0 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o imgen

# For debug reasons
FROM scratch

WORKDIR /root/imagroor

COPY --from=build /src/imgen /root/imagroor/
COPY ./index.html /root/imagroor/
COPY ./img/* /root/imagroor/img/

RUN mkdir /root/imagroor/results

EXPOSE 8080

CMD ["/root/imagroor/imgen"]