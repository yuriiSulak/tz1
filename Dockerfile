FROM ubuntu:18.04
RUN mkdir -p /usr/src/app
COPY app /usr/src/app/app
EXPOSE 8989
CMD ["/usr/src/app/app"]


