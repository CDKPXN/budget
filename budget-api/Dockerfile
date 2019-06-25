FROM scratch
WORKDIR /
ADD ./budget-api /
ADD ./config/config.toml /config/
EXPOSE 80

ENTRYPOINT ["./budget-api"]