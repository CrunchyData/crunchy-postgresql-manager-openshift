#!/bin/bash -x
#
# $1 - openshift URL
# $2 - username
# $3 - password
# $4 - pod template file
#

OPENSHIFT_URL=$1
OPENSHIFT_USERNAME=$2
OPENSHIFT_PASSWORD=$3
TEMPLATE_FILE=$4

/var/cpm/bin/oc login $OPENSHIFT_URL --insecure-skip-tls-verify=true \
	--username=$2 --password=$3
/var/cpm/bin/oc create -f $TEMPLATE_FILE
