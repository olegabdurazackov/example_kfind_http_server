# Sample dockerfile sqlite3
# Пример поискового сервера на Docker
### Usage

    docker image build -t "kf:latest" .
    docker  run --rm  -it -p 127.0.0.1:8080:8080  kf
    localhost:8080/snips
    ...
    ^c
    docker image rm kf:latest
.
