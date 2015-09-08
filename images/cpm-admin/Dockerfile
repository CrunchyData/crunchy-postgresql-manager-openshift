FROM cpm-base-os:latest
MAINTAINER crunchy
ADD  openshift-origin-v1.0.5-96963b6-linux-amd64.tar /var/cpm/bin
ADD bin /var/cpm/bin/
ADD conf /var/cpm/conf/

# the port for the adminapi REST API
EXPOSE 13001

USER postgres

CMD ["/var/cpm/bin/startadmin.sh"]
