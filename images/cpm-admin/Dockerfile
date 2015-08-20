FROM cpm-base-os:latest
MAINTAINER crunchy
ADD openshift-origin-v1.0.3-1695461-linux-amd64.tar /var/cpm/bin/
ADD bin /var/cpm/bin/
ADD conf /var/cpm/conf/

# the port for the adminapi REST API
EXPOSE 13000

VOLUME ["/kubekeys"]

USER postgres

CMD ["/var/cpm/bin/startadmin.sh"]
