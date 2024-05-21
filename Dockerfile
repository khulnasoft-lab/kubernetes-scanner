FROM golang:1.21-bookworm AS build

WORKDIR /home/khulnasoft/src/kubernetes-scanner
COPY . .
RUN go build -o kubernetes-scanner . \
    && chmod 777 kubernetes-scanner \
    && cp /home/khulnasoft/src/kubernetes-scanner/kubernetes-scanner /home/khulnasoft/ \
    && rm -r /home/khulnasoft/src/*

FROM debian:bullseye-slim
MAINTAINER KhulnaSoft Ltd
LABEL khulnasoft.role=system

RUN apt-get update \
    && apt-get install -y bash curl wget git \
    && /bin/sh -c "$(curl -fsSL https://raw.githubusercontent.com/turbot/steampipe/main/install.sh)" \
    && useradd -rm -d /home/khulnasoft -s /bin/bash -g root -G sudo -u 1001 khulnasoft

USER khulnasoft

COPY --from=build /home/khulnasoft/kubernetes-scanner /usr/local/bin/kubernetes-scanner
WORKDIR /opt/steampipe

USER root
ENV VERSION=2.2.0

RUN chown khulnasoft /opt/steampipe /usr/local/bin/kubernetes-scanner

USER khulnasoft
RUN steampipe plugin install steampipe@0.7.0 \
    && steampipe plugin install kubernetes@0.18.1 \
    && git clone https://github.com/turbot/steampipe-mod-kubernetes-compliance.git

ENTRYPOINT ["/usr/local/bin/kubernetes-scanner"]
