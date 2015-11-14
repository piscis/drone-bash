FROM gliderlabs/alpine:3.1
ADD drone-bash /bin/
ENTRYPOINT ["/bin/drone-bash"]
