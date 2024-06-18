package contexts

import (
	"context"

	"github.com/xoctopus/confx/confmws/confmqtt"
	"github.com/xoctopus/x/contextx"

	"github.com/machinefi/sprout-pebble-sequencer/pkg/middlewares/blockchain"
	"github.com/machinefi/sprout-pebble-sequencer/pkg/middlewares/database"
	"github.com/machinefi/sprout-pebble-sequencer/pkg/middlewares/logger"
)

type (
	ctxLogger     struct{}
	ctxMqttBroker struct{}
	ctxMqttClient struct{}
	ctxBlockchain struct{}
	ctxDatabase   struct{}
)

func LoggerFromContext(ctx context.Context) (*logger.Logger, bool) {
	v, ok := ctx.Value(ctxLogger{}).(*logger.Logger)
	return v, ok
}

func WithLoggerContext(v *logger.Logger) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxLogger{}, v)
	}
}

func MqttBrokerFromContext(ctx context.Context) (*confmqtt.Broker, bool) {
	v, ok := ctx.Value(ctxMqttBroker{}).(*confmqtt.Broker)
	return v, ok
}

func WithMqttBrokerContext(v *confmqtt.Broker) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxMqttBroker{}, v)
	}
}

func MqttClientFromContext(ctx context.Context) (*confmqtt.Client, bool) {
	v, ok := ctx.Value(ctxMqttClient{}).(*confmqtt.Client)
	return v, ok
}

func WithMqttClientContext(v *confmqtt.Client) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxMqttClient{}, v)
	}
}

func BlockchainFromContext(ctx context.Context) (*blockchain.Blockchain, bool) {
	v, ok := ctx.Value(ctxBlockchain{}).(*blockchain.Blockchain)
	return v, ok
}

func EthClientFromContextByNetwork(ctx context.Context) (*blockchain.EthClient, bool) {
	v, ok := ctx.Value(ctxBlockchain{}).(*blockchain.Blockchain)
	if !ok {
		return nil, false
	}
	c := v.ClientByNetwork()
	return c, c != nil
}

func WithBlockchainContext(v *blockchain.Blockchain) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxBlockchain{}, v)
	}
}

func DatabaseFromContext(ctx context.Context) (*database.Postgres, bool) {
	v, ok := ctx.Value(ctxDatabase{}).(*database.Postgres)
	return v, ok
}

func WithDatabaseContext(v *database.Postgres) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxDatabase{}, v)
	}
}
