FROM scratch
COPY zoneinfo.tar.gz /
COPY ca-certificates.crt /etc/ssl/certs/
COPY main /
CMD ["/main"]
