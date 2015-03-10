all:
	git clone -b v1.5.0-arm git://github.com/cybercypher/docker.git stem_docker
	(cd stem_docker && make target=arm cross)
clean:
	rm -rf stem_docker
