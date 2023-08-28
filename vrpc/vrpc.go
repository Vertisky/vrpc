package main

import (
	"fmt"

	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Vhost    string
}

type vRPC struct {
	rabbitMQConfig *RabbitMQConfig
	connection     *rabbitmq.Conn
}

func New(config *RabbitMQConfig) (vRPC, error) {
	var err error
	v := vRPC{}
	v.rabbitMQConfig = config

	return v, err
}

func (v *vRPC) connect() error {
	var err error
	v.connection, err = rabbitmq.NewConn(fmt.Sprintf("amqp://%s:%s@%s:%d/%s", v.rabbitMQConfig.Username, v.rabbitMQConfig.Password, v.rabbitMQConfig.Host, v.rabbitMQConfig.Port, v.rabbitMQConfig.Vhost), rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		return err
	}
	return nil
}
