FROM cpm-base-os:latest
MAINTAINER crunchy

ADD bin /var/cpm/bin
ADD conf /var/cpm/conf

USER postgres

CMD ["/var/cpm/bin/start-cpmcontainerapi.sh"]
