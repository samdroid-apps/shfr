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

RUN grep -rl "FORUMS_FILE string = \"forums.json\"" /shfr/forums.go | xargs sed -i 's@forums.json@/data/forums.json@g'
RUN mkdir /data
RUN cp /shfr/forums.json /data/forums.json
VOLUME /data

RUN cd /shfr; go build
ENTRYPOINT ["/shfr/shfr"]
