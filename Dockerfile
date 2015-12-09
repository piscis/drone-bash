FROM python:2.7
RUN pip install fabric
ADD drone-bash /bin/
ENTRYPOINT ["/bin/drone-bash"]
