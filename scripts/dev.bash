cd ..
docker build ./ -t=devbox
winpty docker run -it -p 8080:8080 -v Go\src\github.com\mccjul\helloworld\app devbox