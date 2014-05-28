FROM flynn/busybox
MAINTAINER Charlie Knudsen <charlie.knudsen@gmail.com>

ADD ./build/linux/docker-testing /bin/docker-testing

EXPOSE 9090

ENTRYPOINT ["/bin/docker-testing"]
CMD []