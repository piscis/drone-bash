FROM python:2.7
RUN pip install fabric
ADD drone-fabric /bin/
ENTRYPOINT ["/bin/drone-fabric"]
