FROM ubuntu:14.04

RUN apt-get update
RUN apt-get install git golang -y
RUN git clone https://github.com/samdroid-apps/shfr /shfr

RUN mkdir /go
ENV GOPATH /go

RUN git clone https://github.com/gin-gonic/gin /go/src/github.com/gin-gonic/gin/
RUN git config --global user.email "you@example.com"
RUN git config --global user.name "Your Name"
RUN cd /go/src/github.com/gin-gonic/gin/; git fetch origin app_engine; git merge origin/app_engine -m "Yeah"

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/andreaskoch/go-fswatch

RUN grep -rl "RECORDS_FILE string = \"records.json\"" /shfr/record.go | xargs sed -i 's@records.json@/data/records.json@g'
RUN mkdir /data
VOLUME /data

RUN cd /shfr; go build

EXPOSE 8000
ENTRYPOINT ["/shfr/shfr"]
