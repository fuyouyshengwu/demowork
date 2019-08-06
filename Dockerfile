FROM registry.icp.com:5000/cks/cks-ubuntu:test
ADD main /demo/work
RUN chmod +x /demo/work
EXPOSE 80
WORKDIR /demo
ENTRYPOINT ["bash","-c"]
CMD ["/demo/work"]
