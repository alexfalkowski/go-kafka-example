package kafka_test

import (
	"testing"

	"github.com/Shopify/sarama"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKafka(t *testing.T) {
	Convey("Given I have a valid kafka setup", t, func() {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 5
		config.Consumer.Return.Errors = true

		addrs := []string{"localhost:9092"}
		topic := "test"

		producer, err := sarama.NewSyncProducer(addrs, config)
		So(err, ShouldBeNil)

		defer producer.Close()

		consumer, err := sarama.NewConsumer(addrs, config)
		So(err, ShouldBeNil)

		defer consumer.Close()

		par, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
		So(err, ShouldBeNil)

		Convey("When I send a message", func() {
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(topic),
			}

			partition, offset, err := producer.SendMessage(msg)
			So(err, ShouldBeNil)
			So(offset, ShouldBeGreaterThanOrEqualTo, 0)
			So(partition, ShouldBeGreaterThanOrEqualTo, 0)

			Convey("Then I should receive a message", func() {
				msg := <-par.Messages()

				So(msg.Value, ShouldResemble, []byte(topic))
			})
		})
	})
}
