FROM ubuntu:14.04
RUN apt-get update && apt-get install git build-essential libssl-dev
ADD drone-bash /bin/
ENTRYPOINT ["/bin/drone-bash"]
