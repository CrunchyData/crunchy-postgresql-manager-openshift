FROM cpm-base-os:latest
MAINTAINER crunchy

# open up the backupserver port
EXPOSE 13000

ADD bin /var/cpm/bin
ADD conf /var/cpm/conf

USER root

CMD ["/var/cpm/bin/start-backupserver.sh"]
