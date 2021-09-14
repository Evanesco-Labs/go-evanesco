FROM centos:centos7.9.2009
ADD . /go-evanesco
WORKDIR /go-evanesco
RUN mkdir ./avisnode && cp ./build/bin/eva ./avisnode && cp ./verifykey.txt ./avisnode && cp ./avis.json ./avisnode && cp ./avis.toml ./avisnode
WORKDIR /go-evanesco/avisnode
RUN mkdir data
CMD ["/bin/bash"]
