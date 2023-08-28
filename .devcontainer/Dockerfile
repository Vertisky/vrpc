FROM etma/devcontainer-go:1.21.0-ubuntu-kube

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    docker-compose \
    2>&1

COPY setup-golang.sh /root
RUN chmod +x /root/setup-golang.sh
