## How to use ssh / mongo shell

Run the following commands.
Access the mongo shell through the docker container:

```
docker exec -it mongodb bash
```

Run the mongo shell with privileges:

```
mongo --username root --password password --authenticationDatabase admin
```

Show the dbs:

```
db
```

Use the database - if none exists it will create one with the name:

```
use "nameofdatabase"
```

Insert an item into a new collection called myNewCollection - it will auto create:

```
db.myNewCollection.insertOne( {x:1});
```

Show all the collections:

```
show collections
```

Show all the times in the collection:

```
db.myNewCollection.find()
```