FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/roykachouh/fileWritingService/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go get ./...; go build -o myapp; cp myapp /app/

ENTRYPOINT ["./myapp"]
