FROM nginx:1.19.10

# WORKDIR é um comando que define o diretório de trabalho para os comandos RUN, CMD, ENTRYPOINT, COPY e ADD
WORKDIR /app

RUN apt-get update && \
    apt-get install vim -y

# COPY é um comando que copia arquivos ou diretórios do host para o container
COPY html /usr/share/nginx/html