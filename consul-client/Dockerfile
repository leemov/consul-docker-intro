FROM ubuntu:18.04
WORKDIR /app
# install pre-requisite packages
RUN apt-get update && \
      apt-get -y install sudo 
RUN apt-get install -y curl
RUN apt-get update && apt-get install -y gnupg2
RUN apt-get update && apt-get install -y lsb-release && apt-get clean all
RUN apt-get install -y software-properties-common

# from consul download manual
RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
RUN sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
RUN sudo apt-get update && sudo apt-get install consul

RUN sudo apt install -y dnsutils
RUN sudo apt install -y telnet
RUN apt-get install -y iputils-ping
RUN sudo apt-get install -y iptables

COPY . .
RUN chmod 775 script.sh