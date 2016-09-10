# Bidu

Cuidado com o bidu, ele pega seu piru.

## Unit Tests

To run `unit tests` type:

```
make test
```

## Database

Here I will show how to run postgres on Docker, but you can install and run
postgres in your machine.

```
make docker
```

Get the id of the running container:

```
sudo docker ps
```

Shell into the container:

```
sudo docker exec -it <container-id> /bin/bash
```

Run psql:

```
exec psql -U postgres
```

Create the database:

```
create database bidu;
```

## Database Migrations

After creating the database we can populate it with the following command:

```
make migrate
```

## Dredd

Dredd is a language agnostic command-line tool for testing API documentation written in
the API Blueprint format against its backend implementation.

Install Dredd:

```
# npm install -g dredd
```

Run the app:

```
$ make run
```

Run Dredd:

```
$ make dredd
```

output:

```
info: Beginning Dredd testing...
pass: POST /api/sites duration: 42ms
pass: GET /api/sites/4DTest duration: 18ms
pass: PUT /api/sites/4DTest duration: 10ms
pass: DELETE /api/sites/5DTest duration: 11ms
complete: 4 passing, 0 failing, 0 errors, 0 skipped, 4 total
complete: Tests took 87ms
```
