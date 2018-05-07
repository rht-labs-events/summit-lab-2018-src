FROM scratch
MAINTAINER Jessica Forrester <jforrest@redhat.com>
COPY bin/hello-openshift /hello-openshift
EXPOSE 8080 8888
RUN chmod 777 /hello-openshift
USER 1001
ENTRYPOINT ["/helo-openshift"]
