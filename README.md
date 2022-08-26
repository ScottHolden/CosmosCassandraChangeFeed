# CosmosCassandraChangeFeed
A small demo using the CosmosDB Change Feed within the Cassandra API in Go

## Getting started
1. Make sure you have the [Go SDK installed](https://go.dev/dl/).  
   _Note: This was developed with Go 1.16_
2. Clone this repo or use the "Download ZIP" button under the "<> Code" button above.
3. Create a Cosmos DB Account using the Cassandra API, you can either deploy via the [Azure Portal](https://ms.portal.azure.com/#create/Microsoft.DocumentDB) or [Azure CLI](https://docs.microsoft.com/en-us/azure/cosmos-db/scripts/cli/cassandra/create).  
   _Note: The keyspace and table will be automatically created, you just need to create the account._
4. Modify the "config.yml" file and update the **Host**, **Port**, **Username**, and **Password** values to point towards your CosmosDB.
5. Open a terminal pointed to the folder you cloned (or extracted) this repo to, and execute `go run .`