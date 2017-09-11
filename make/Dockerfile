FROM centurylink/ca-certs
EXPOSE 5600

COPY release/linux/amd64/ymir /ymir

ENTRYPOINT ["/malcolm"]
CMD ["server"]
