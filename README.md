# RedTmon
Simple but effective RedT health monitor, providing the activity of Validator nodes.

## Running with Docker

The easiest way is to run the container. You can just use the image `hesusruiz\redtmon` in Dockerhub.

Or you can build the image yourself if you want, using the Dockerfile in this repository.

You can also use the `compose.yaml` as an example for running a container with Docker Compose.

## Building the container

```
docker build -t hesusruiz/redtmon:v0.4.3 .
```
