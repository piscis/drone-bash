FROM ubuntu:14.04
RUN apt-get update
ADD drone-bash /bin/
#ENTRYPOINT ["/bin/drone-bash"]
