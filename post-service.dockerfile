FROM alpine:latest
RUN mkdir /app

COPY ./post-service.bin /app

# Run the server executable
CMD [ "/app/post-service.bin" ]