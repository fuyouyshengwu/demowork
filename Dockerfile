FROM registry.inspurcloud.cn/picp_inspurtest14/cus-go:test
ADD main /demo/work
RUN chmod +x /demo/work
EXPOSE 80
WORKDIR /demo
ENTRYPOINT ["bash","-c"]
CMD ["/demo/work"]
