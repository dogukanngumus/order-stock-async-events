using order_service.Messaging;
using order_service.Models;

namespace order_service.Services;

public class OrderService
{
    private readonly RabbitMqPublisher _publisher;
    private static readonly List<Order> _orders = new();

    public OrderService(RabbitMqPublisher publisher)
    {
        _publisher = publisher;
    }


    public Order CreateOrder(Order order)
    {
        _orders.Add(order);
        _publisher.Publish(order); 
        return order;
    }

    public IEnumerable<Order> GetAllOrders() => _orders;
}
