# greenleaf - simple query builder for MongoDB

[![godoc](https://godoc.org/github.com/slavabobik/greenleaf?status.png)](https://godoc.org/github.com/slavabobik/greenleaf)
    

## Installation
To install use:

```bash
 go get github.com/slavabobik/greenleaf
```   


## Quick examples

```golang
package main

import (
	"context"
	"fmt"

	"github.com/slavabobik/greenleaf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.TODO()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("testing").Collection("test")
	collection.InsertOne(ctx, bson.M{"pet": "dog"})

	filter := greenleaf.
		Filter().
		Eq("pet", "dog").
		Build()

	result := collection.FindOne(ctx, filter)
	var document bson.M
	result.Decode(&document)

	fmt.Print(document, "xxx")
}

```