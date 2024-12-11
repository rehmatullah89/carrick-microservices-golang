package queue

import (
	"carrick-js-api/pkgs/config"
	"carrick-js-api/pkgs/logger"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"net/url"
	"sync"
)

type rabbitMQSingleton struct {
	connect *amqp.Connection
}

var (
	rabbitMQOnce     sync.Once
	rabbitMQInstance *rabbitMQSingleton
)

type NewUrlTask struct {
	Publisher_Id      uint
	Traffic_Source_Id uint
	Url_Path          string
}

func GetRabbitMQInstance() *rabbitMQSingleton {
	logger := logger.GetLoggerInstance()

	if rabbitMQInstance == nil {
		rabbitMQOnce.Do(
			func() {
				rabbitMQInstance = &rabbitMQSingleton{}

				if _, err := rabbitMQInstance.ReConnect(); err != nil {
					logger.Error(err)
				}
			})
	}

	return rabbitMQInstance
}

func (c *rabbitMQSingleton) Close() {
	c.connect.Close()
}

func (c *rabbitMQSingleton) ReConnect() (bool, error) {
	logger := logger.GetLoggerInstance()

	if c.connect == nil || c.connect.IsClosed() {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
			url.QueryEscape(config.AppConfig.RabbitMQ.User),
			url.QueryEscape(config.AppConfig.RabbitMQ.Password),
			config.AppConfig.RabbitMQ.Host,
			config.AppConfig.RabbitMQ.Port))
		if err == nil {
			logger.Info("RabbitMQ connected")
			c.connect = conn

			return true, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (c *rabbitMQSingleton) PushMessage(topicName string, task NewUrlTask) {
	logger := logger.GetLoggerInstance()

	if _, err := c.ReConnect(); err != nil {
		logger.Error(err)

		return
	}

	ch, err := c.connect.Channel()
	if err != nil {
		logger.Error(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		topicName, // name
		true,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.Error(err)
		return
	}

	m, err := json.Marshal(task)
	if err != nil {
		logger.Error(err)
		return
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        m,
			DeliveryMode: 2,
		})
	if err != nil {
		logger.Error(err)
		return
	}
}
