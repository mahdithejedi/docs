# important docker commands

```docker run = docker create + docker start```

### point: exactly docker run is equal with ```dokcer run = docker create + docker start --attach(=-a)```

### point: port mapping in docker ```docker run -p <INCOMMING_LOCALHOST_PORT>:<DOCKER_INCOMMING_PORT> ...``` eg: docker run -p 8080:8080 capitanos/node

<br />

```docker build -t<TAGNAME> <DOCKERPATH>```    this command will build a docker image from *Dockerfile*

#### point 1 : the format of the tag name should follow `YourDockerID / Repo/projectName : Version` for example seyedmm021/redis:lastest 

#### point 2 : you can execute or do any stuff with the tagged docker without specify the docker version eg: `docker run seyedmm021/redis`

#### point 3 : you can rub specific Docker with flaf `-f` like ```docker build -f /PATH/TO/DOCKERFILE```
<br />

```docker ps``` show current docker proccess with `-all` or `-a` you can see all the process even which is done
<br />

```docker system prune``` you can remove uneccessry images and containers
<br />

```docker exect -it <DOCKER_ID> <COMMEND>``` for run a comment into a docker , the `-i`  means accept input and `-t` means show output
<br />

```doker commit -c <COMMAND> <DOCKER_CONTAINER_ID>``` create a image out of a container eg:`docker commit -c 'CMD ["redis-server"]' ef2313dc` which *ef2313dc* is a container
<br />

# exactly what is a docker continer or image?

***prequires***

[*what is a linux container*](https://opensource.com/resources/what-are-linux-containers)

[*what is a ps in linux*](https://www.geeksforgeeks.org/ps-command-in-linux-with-examples/)

**more focus in linux containers**

[*part one*](https://www.linuxjournal.com/content/everything-you-need-know-about-linux-containers-part-i-linux-control-groups-and-process)

[*part two*](https://www.linuxjournal.com/content/everything-you-need-know-about-linux-containers-part-ii-working-linux-containers-lxc)

***docker image***

[*stackoverflow answer*](https://stackoverflow.com/questions/27359771/whats-inside-a-docker-image-container)

[*Docker image*](https://searchitoperations.techtarget.com/definition/Docker-image)

[*image vs container*](https://stackify.com/docker-image-vs-container-everything-you-need-to-know/)

<br />
<hr />

# Docker-compose

`docker-compuse up` == docker run image

    `docker-compose up --build` == docker build . + docker run image


opposit of **docker-compuse** up we can use 

    `docker-compose down` to down group of dockers 

<hr />

## <font color='yellow'>docker restart policy [yaml]</font>

**Directive:`restart`**

docker-compose decision for restarting or not


    "no"/always/on-failure/unless-stopped

<hr />

## <font color='#f8615a'>docker process status(ps) [Cli Commend]</font>

**CliCommend:`ps`**

excatly like docker ps here you can see the status of docker-compose

#### point:need docker-compose.yml file


# [Docker volumes](https://docs.docker.com/storage/volumes/)

map from a folder inside of a container to a folder outside of a container the syntacs is like `DOCKER RUN -v /PATH/TO/FOLDER` or `DOCKER RUN --volume /PATH/TO/FOLDER` 

Example: ```DOCKER RUN -p 3000:3000 -v /app/node_modules -v $(pwd):/app <IMAGE_ID>```

exually the **$(pwd):/app** means map /app which is inside of a container to a folder out of container in this exmple which is $(pwd)  although the first -v point to inside the container and you may think it is not usefull but we use that to say :'hey docker please do not override the access of /app for node_modules


<hr />

## <font color='lightgreen'>FROM statement [Dockerfile]</font>

eg1: `FROM nginx:latest as nginx`

#### point1: you can use the FROM command with copy eg:`COPY --from=nginx /NGINX/DOCKER/PATH /DESTENATION/PATH
