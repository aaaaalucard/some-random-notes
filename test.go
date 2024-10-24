
import (
  "context"
  "fmt"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
  // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  //opts := options.Client().ApplyURI("mongodb+srv://hankyuan2:D1fficultP%40ssw0rd@clusterabisf.kcskq.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  //opts := options.Client().ApplyURI("mongodb+srv://hankyuan2:P@ssw0rd@clusterabisf.kcskq.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  opts := options.Client().ApplyURI("mongodb+srv://test:<db_password>@dedicatedcluster-172455.ukz8e.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  //opts := options.Client().ApplyURI("mongodb+srv://asdfasdf:******@dedicatedcluster-172455.kcskq.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  //opts := options.Client().ApplyURI("mongodb+srv://hankyuan:<password>@cluster0.kcskq.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  //opts := options.Client().ApplyURI("mongodb+srv://hankyuan:99996666@clusterabisf.kcskq.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  //opts := options.Client().ApplyURI("mongodb+srv://hankyuan2:******@clusterabisf.kcskq.mongodb.net/?retryWrites=true&w=majority&appName=DedicatedCluster-1724559993").SetServerAPIOptions(serverAPI)
  // Create a new client and connect to the server
  client, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  // Send a ping to confirm a successful connection
  if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
    panic(err)
  }
  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
