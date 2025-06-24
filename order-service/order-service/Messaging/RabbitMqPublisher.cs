using System.Text;
using System.Text.Json;
using order_service.Models;
using RabbitMQ.Client;

namespace order_service.Messaging;

public class RabbitMqPublisher
{
    private readonly IConnection _connection;
    private readonly IModel _channel;

    public RabbitMqPublisher(IConfiguration configuration)
    {
        var factory = new ConnectionFactory
        {
            HostName = configuration["RabbitMq:Host"] ?? "localhost"
        };

        _connection = factory.CreateConnection();
        _channel = _connection.CreateModel();
        _channel.QueueDeclare(queue: "stock-update", durable: false, exclusive: false, autoDelete: false);
    }

    public void Publish(Order order)
    {
        var message = JsonSerializer.Serialize(order);
        var body = Encoding.UTF8.GetBytes(message);
        _channel.BasicPublish(exchange: "", routingKey: "stock-update", body: body);
    }
}
