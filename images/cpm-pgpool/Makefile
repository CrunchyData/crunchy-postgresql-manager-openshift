OSFLAVOR=centos7

all:
	rm -rf bin
	mkdir bin
	cp ../../sbin/*  ./bin
	cp ../../../../../../bin/cpmcontainerapi  ./bin
	sudo docker build -t cpm-pgpool-os -f Dockerfile.$(OSFLAVOR) .
	sudo docker tag -f cpm-pgpool-os:latest crunchydata/cpm-pgpool-os
