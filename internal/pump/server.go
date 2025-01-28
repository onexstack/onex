package pump

import (
	"context"
	"time"

	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"

	kafkaconnector "github.com/onexstack/onex/pkg/streams/connector/kafka"
	mongoconnector "github.com/onexstack/onex/pkg/streams/connector/mongo"
	"github.com/onexstack/onex/pkg/streams/flow"
)

// Config contains application-related configurations.
type Config struct {
	KafkaOptions *genericoptions.KafkaOptions
	MongoOptions *genericoptions.MongoOptions
}

// Server represents the web server.
type Server struct {
	config  kafka.ReaderConfig
	colName string
	db      *mongo.Database
}

// addUTC appends a UTC timestamp to the beginning of the message value.
var addUTC = func(msg kafka.Message) kafka.Message {
	timestamp := time.Now().Format(time.DateTime)

	// Concatenate the UTC timestamp with msg.Value
	msg.Value = []byte(timestamp + " " + string(msg.Value))
	return msg
}

// NewServer initializes and returns a new Server instance.
func (cfg *Config) NewServer(ctx context.Context) (*Server, error) {
	client, err := cfg.MongoOptions.NewClient()
	if err != nil {
		return nil, err
	}

	server := &Server{
		config: kafka.ReaderConfig{
			Brokers:           cfg.KafkaOptions.Brokers,
			Topic:             cfg.KafkaOptions.Topic,
			GroupID:           cfg.KafkaOptions.ReaderOptions.GroupID,
			QueueCapacity:     cfg.KafkaOptions.ReaderOptions.QueueCapacity,
			MinBytes:          cfg.KafkaOptions.ReaderOptions.MinBytes,
			MaxBytes:          cfg.KafkaOptions.ReaderOptions.MaxBytes,
			MaxWait:           cfg.KafkaOptions.ReaderOptions.MaxWait,
			ReadBatchTimeout:  cfg.KafkaOptions.ReaderOptions.ReadBatchTimeout,
			HeartbeatInterval: cfg.KafkaOptions.ReaderOptions.HeartbeatInterval,
			CommitInterval:    cfg.KafkaOptions.ReaderOptions.CommitInterval,
			RebalanceTimeout:  cfg.KafkaOptions.ReaderOptions.RebalanceTimeout,
			StartOffset:       cfg.KafkaOptions.ReaderOptions.StartOffset,
			MaxAttempts:       cfg.KafkaOptions.ReaderOptions.MaxAttempts,
		},
		colName: cfg.MongoOptions.Collection,
		db:      client.Database(cfg.MongoOptions.Database),
	}

	return server, nil

}

// Run starts the server and listens for termination signals.
// It gracefully shuts down the server upon receiving a termination signal.
func (s *Server) Run(ctx context.Context) error {
	source, err := kafkaconnector.NewKafkaSource(ctx, s.config)
	if err != nil {
		return err
	}

	filter := flow.NewMap(addUTC, 1)

	sink, err := mongoconnector.NewMongoSink(ctx, s.db, mongoconnector.SinkConfig{
		CollectionName:            s.colName,
		CollectionCapMaxDocuments: 2000,
		CollectionCapMaxSizeBytes: 5 * genericoptions.GiB,
		CollectionCapEnable:       true,
	})
	if err != nil {
		return err
	}

	log.Infof("Successfully start pump server")
	source.Via(filter).To(sink)

	return nil

}
