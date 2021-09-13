FROM centos:centos7.9.2009

RUN yum update -y && yum install epel-release zip unzip wget curl -y && yum install golang -y

WORKDIR /opt/gopath/src/github.com/evanesco/
ADD ./miner-linux.zip /opt/gopath/src/github.com/evanesco/
ADD ./QmNpJg4jDFE4LMNvZUzysZ2Ghvo4UJFcsjguYcx4dTfwKx /opt/gopath/src/github.com/evanesco/
RUN unzip miner-linux.zip && mv ./miner-linux miner && rm miner-linux.zip && mv ./QmNpJg4jDFE4LMNvZUzysZ2Ghvo4UJFcsjguYcx4dTfwKx ./miner

CMD ["/bin/bash"]
