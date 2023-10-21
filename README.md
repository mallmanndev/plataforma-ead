# PLATAFORMA EAD

## 📃 SOBRE

Projeto de uma plataforma de ensino a distância(EAD) utilizando microsserviços em GO e também Next 13.

## 🔧 TOOLS

O projeto está sendo desenvolvido utilizando as seguintes tecnologias

- [NextJS](https://nextjs.org/)
- [GO](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [MongoDB](https://www.mongodb.com/)
- [gRPC](https://grpc.io/)

## SERVICES

O projeto vai contar com dois microsserviços e uma BFF/Front

### Arquitetura do projeto
![Diagrama da arquitetura](architecture.png)

### BFF/Front

Vamos utilizar o NextJS como BFF e tembém como front-end da aplicação.

### Service Core

Serviço principal da plataforma.

🔧 Techs:

- GO
- PostgresSQL

🎯 Responsabilidades:

- Cadastro e alteração de usuários
- Autorização e autenticação de usuarios
- Gerenciamento de usuarios da plataforma
- Gerenciamento de matriculas de curso
- Avaliação de atividades

### Service course

Serviço para o gerenciamento e disponibilização de cursos.

🔧 Techs:

- GO
- Banco de dados MongoDB

🎯 Responsabilidades:

- Cadastro e alteração de cursos
- Upload de arquivos
- Streaming de videos
- Gerenciamento de progresso de curso

// ffmpeg -i video-editado.mp4 -c:v libx264 -vf scale=640:480 -f segment -segment_time 10 -segment_list test_480/segment-list.m3u8 -c copy -re
set_timestamps 1 test_480/segmento%d.mp4

Redimensiona o video mantendo a qualidade

ffmpeg \
    -i input.mp4 \
    -vf scale=1280:720 \
    -preset slow \
    -crf 18 \
    -hls_time 10 \
    -hls_list_size 0 \
    -hls_segment_filename "resolucao-720/segment%d.ts" \
    -f hls resolucao-720/index.m3u8

ffmpeg \
-i input.mp4 \
-vf scale=640:480 \
-preset slow \
-crf 18 \
-hls_time 10 \
-hls_list_size 0 \
-hls_segment_filename "resolucao-480/segment%d.ts" \
-f hls resolucao-480/index.m3u8

Somente gerar segmentos:
ffmpeg \
    -i input.mp4 \
    -hls_time 10 \
    -hls_list_size 0 \
    -hls_segment_filename "resolucao-1080/segment%d.ts" \
    -f hls resolucao-1080/index.m3u8

curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | \
  sudo apt-key add -
distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | \
  sudo tee /etc/apt/sources.list.d/nvidia-docker.list
sudo apt-get update

# Install nvidia-docker2 and reload the Docker daemon configuration
sudo apt-get install -y nvidia-docker2
sudo pkill -SIGHUP dockerd