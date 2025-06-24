using System.Text.Json;
using Microsoft.EntityFrameworkCore;
using order_service.Messaging;
using order_service.Models;

namespace order_service.Services;

public class OrderService
{
    private readonly AppDbContext _context;

    public OrderService(AppDbContext context)
    {
        _context = context;
    }

    public async Task<Order> CreateOrderAsync(Order order)
    {
        await _context.Orders.AddAsync(order);

        var outboxMessage = new OutboxMessage
        {
            EventType = "OrderCreated",
            Payload = JsonSerializer.Serialize(order)
        };
        await _context.OutboxMessages.AddAsync(outboxMessage);

        await _context.SaveChangesAsync();

        return order;
    }
    public async Task<IEnumerable<Order>> GetAllOrdersAsync()
    {
        return await _context.Orders.ToListAsync();
    }
}