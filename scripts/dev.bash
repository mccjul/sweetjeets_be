docker build ./
docker run -it -p 8080:8080 -v [put your path here]/app:/go/src/github.com/user/myProject/app [image id]