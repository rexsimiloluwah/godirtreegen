FROM golang:1.16-alpine 

WORKDIR /app 

ARG MOUNT_DIR=/host
RUN mkdir -p ${MOUNT_DIR}}

RUN mkdir -p cmd 
COPY cmd/treegen.go cmd
COPY run.sh main.go .folderignore go.mod go.sum /app/ 

RUN ls -la

RUN go mod download 
# build 
RUN go build -o godirgen 
COPY run.sh ./
RUN chmod +x ./run.sh 
ENTRYPOINT ["./run.sh"]
# VOLUME [${MOUNT_DIR}}]
