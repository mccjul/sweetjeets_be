docker build ./ --build-arg app_env=production
docker run -i -t -p 8080:8080 [image id]