FROM scratch

COPY test /test

ENTRYPOINT ["/test"]
