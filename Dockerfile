FROM keinos/sqlite3
LABEL maintainer="olab@email.su"
WORKDIR .
COPY . .
EXPOSE 8080
CMD ["./server_http_find_snip"]
VOLUME /db

