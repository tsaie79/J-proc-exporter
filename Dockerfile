ARG BASE_IMAGE=debian:bullseye-slim
FROM $BASE_IMAGE as base

# create a stage for building the process-exporter binary
FROM --platform=$BUILDPLATFORM golang:1.17 AS process_exporter_build
ARG PROCESS_EXPORTER_REPO_URL="https://github.com/ncabatoff/process-exporter.git"
RUN git clone -b master "${PROCESS_EXPORTER_REPO_URL}" /go/src/github.com/ncabatoff/process-exporter
ARG TARGETARCH 
ARG BUILDPLATFORM 
WORKDIR /go/src/github.com/ncabatoff/process-exporter
ADD . .
# Build the process-exporter command inside the container.
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH make build


# create a stage that compile the main.go file
FROM --platform=$BUILDPLATFORM golang:1.21 AS gen_config_build
ARG TARGETARCH
ARG BUILDPLATFORM
COPY ./generate /generate
WORKDIR /generate
RUN go build -o /gen-config main.go


FROM base
COPY --from=process_exporter_build /go/src/github.com/ncabatoff/process-exporter/process-exporter /usr/local/bin/process-exporter
COPY --from=gen_config_build /gen-config /usr/local/bin/gen-config

COPY run.sh /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/run.sh

RUN apt-get update && apt-get install procps -y 


ENTRYPOINT ["/usr/local/bin/run.sh"]

#------------------ Above is from procee-exporter ------------------