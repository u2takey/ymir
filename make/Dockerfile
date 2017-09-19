FROM centurylink/ca-certs
EXPOSE 5600

COPY release/linux/amd64/ymir /make/release/ymir
COPY index.html /make/index.html 

ENTRYPOINT ["/make/release/ymir"]
CMD ["server"]
