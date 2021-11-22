# Wildfire Go API

This application is intended to showcase my ability to create a resilient
golang api. Nginx is used a load balancer and docker is used for container
managment and creation. Nginx was selected for two reasons. Firstly, nginx is a
battle tested technology. It's reliability and speed are well documented. 
Docker was chosen for its ubuquity. Docker is supported by all major cloud 
solutions providers and well known for it's security. 

This project is provided with two goals in mind. First, it should be easy to run
locally and test. Second, the api is written in a way that one can easily deploy
to a single server or simply use the associated Dockerfile to create a 
standalone image that can be inserted to any number of cloud 
solutions/architecture. Nginx would likely be replaced by your cloud provider's 
own infrastructure for load balancing and reverse proxy. 

##How to Install Your Project

You will need Docker 20.10.8 or higher and Docker Compose 1.25 or higher. 

- Instructions for Docker installation can be found at https://docs.docker.com/get-docker/
- Instructions for Docker Compose installation can be found at https://docs.docker.com/compose/install/

Once these are installed go to the local instance of this github repo and run
the following:

`docker-compose --compatibility up --build`

That's it. You're ready use curl or your favorite web browser 

`curl localhost:80/api/v0/joke`

OR

1. Go to web browser of choice
2. Type in `localhost:80/api/v0/joke`
3. Refresh the page for new jokes

## MIT License

Copyright (c) [2021] [Mac Farnsworth]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
