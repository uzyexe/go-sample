FROM golang:1.6.3-onbuild

EXPOSE 8080
CMD ["app", "-s"]
