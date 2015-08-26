
## Configure Origin container as documented at github.com/openshift/origin

## Configure host network
edit /etc/resolv.conf
edit /etc/sysconfig/docker

add static IP address into each that we use for DNS resolution

## the Origin security context to allow RunAsAny and allow HostDir volume plugin

~~~~~~~~
sudo docker exec -it origin bash
oc edit scc restricted
~~~~~~~~

## Set the selinux label correctly

~~~~~~~~
chcon sandbox.... /var/lib/openshift on the host 
~~~~~~~~

## start cpm-admin

~~~~
oc process -f cpm-admin.json | oc create -f -
~~~~



