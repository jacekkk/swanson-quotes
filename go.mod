module swanson

go 1.14

require (
	// cloud.google.com/go v0.56.0
	// cloud.google.com/go/bigquery v1.4.0
	// cloud.google.com/go/pubsub v1.2.0
	// cloud.google.com/go/storage v1.6.0
	// github.com/envoyproxy/go-control-plane v0.9.4
	// github.com/go-sql-driver/mysql v1.5.0
	// google.golang.org/grpc v1.28.1
	github.com/GoogleCloudPlatform/cloudsql-proxy v1.17.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.7.4
)

replace cloud.google.com/go/bigquery => cloud.google.com/go/bigquery v1.4.0

replace cloud.google.com/go/pubsub => cloud.google.com/go/pubsub v1.2.0

replace google.golang.org/grpc => google.golang.org/grpc v1.28.1
