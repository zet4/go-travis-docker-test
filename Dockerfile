FROM scratch
ADD main /
EXPOSE 3333
CMD ["/main"]