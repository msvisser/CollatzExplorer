FROM scratch

COPY ./collatz /

EXPOSE 8000
CMD ["/collatz"]
