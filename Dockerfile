FROM scratch

ADD ./bin/demo /bin/demo

ENTRYPOINT ["/bin/demo"]
