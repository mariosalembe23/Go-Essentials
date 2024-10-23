# Usar a imagem oficial do Go como base
FROM golang:1.20

WORKDIR /golang

COPY . .


CMD ["/bin/bash"]