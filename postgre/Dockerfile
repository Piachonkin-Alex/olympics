FROM bitnami/postgresql:12.4.0

USER root
RUN apt-get update && apt-get -y install python3-pip libpq-dev jq
RUN pip3 install yandex-pgmigrate
