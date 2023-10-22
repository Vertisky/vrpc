package main

import (
	"fmt"

	"github.com/vertisky/vrpc/service"
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
	baseName       string
	rabbitMQConfig *RabbitMQConfig
	connection     *rabbitmq.Conn
	services       []service.Service
}

func New(baseName string, config *RabbitMQConfig) (vRPC, error) {
	var err error
	v := vRPC{
		baseName: baseName,
	}
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

func (v *vRPC) Close() error {
	return v.connection.Close()
}

func (v *vRPC) AddService(name string, service *service.Service) error {
	v.services = append(v.services, *service)
	return nil
}
