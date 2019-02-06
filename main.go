package main

import "github.com/mongodb/mongo-go-driver/mongo"

func main {
  connectionStr := os.Getenv("DefaultConnection")
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  client, err := mongo.Connect(ctx, connectionStr)
}
