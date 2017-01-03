FROM scratch
ADD counter-test /
EXPOSE 3333
CMD ["/counter-test"]