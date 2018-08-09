cd ..
docker build ./ --build-arg app_env=production -t=prodbox
winpty docker run -i -t -p 8080:8080 prodbox