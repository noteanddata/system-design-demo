
# about
a simple tinyrul demo provided by [http://www.noteanddata.com](http://www.noteanddata.com)

# how to run 

## run in docker 
1. build docker image
```
docker build . -t tinyurl
```
2. run docker container 
```
docker run -dit -p 8080:8080 tinyurl
```
3. visit tinyurl web page 
http://localhost:8080

## run locally 
1. run db/db-init.sql in local MySQL database
2. go run tinyurl.go

# todo list 
- write a docker file and run inside a container
- add better error handling and refactor the codes
- benchmark the performance

