# mdb2json

### Objective

The purpose of this repository is to get successfully data from the .bok files using mdb tools and eventually in json format ready to be indexed in a NoSQL database like Elasticsearch or MongoDB.

The following steps are to be taken by me to achieve the above goal.

1.  Successfully use mdb-tools to export queries and through sqlite3
2.  Using goroutines decrease the time it takes to export 1000's of books.


### checklist

- [ ]   Dump schema into Sql Database
- [ ]   Dump data into Sql Database
- [ ]   refactor the code, so sql specific functions are in their own go file
- [ ]   Work out the schema / structure in json
- [ ]   Extract Database tables and store as json objects in json file
