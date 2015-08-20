OSFLAVOR=centos7

all:
	        sudo docker build -t cpm-os -f Dockerfile.$(OSFLAVOR) .
		sudo docker tag -f cpm-os:latest crunchydata/cpm-os
