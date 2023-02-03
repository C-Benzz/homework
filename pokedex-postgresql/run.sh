docker rmi pokedex-app:1.0
docker build -t pokedex-app:1.0 .
# docker run -d --net=bridge pokedex-app:1.0
# docker run -d pokedex-app:1.0
# docker rmi pokedex-app:1.0