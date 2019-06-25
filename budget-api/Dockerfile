FROM busybox
WORKDIR /home
ADD ./budget-api /home/budget-api
ADD ./config/config.toml /home/config.toml
EXPOSE 80
# USER root
RUN chmod +x /home/budget-api
RUN ls -l /home/
ENTRYPOINT ["/home/budget-api"]