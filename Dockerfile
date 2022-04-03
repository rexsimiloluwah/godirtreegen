FROM golang:1.16-alpine 

WORKDIR /app 

ARG MOUNT_DIR=/host
RUN mkdir -p ${MOUNT_DIR}}

COPY . ./ 
RUN go mod download 
# build 
RUN go build -o godirgen 
COPY run.sh ./
RUN chmod +x ./run.sh 
ENTRYPOINT ["./run.sh"]
# VOLUME [${MOUNT_DIR}}]
