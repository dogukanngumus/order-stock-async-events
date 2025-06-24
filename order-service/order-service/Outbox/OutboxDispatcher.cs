using System;
using Microsoft.EntityFrameworkCore;
using order_service.Messaging;
using order_service.Models;

namespace order_service.Outbox;

public class OutboxDispatcher
{
    private readonly AppDbContext _context;
    private readonly RabbitMqPublisher _publisher;
    private readonly ILogger<OutboxDispatcher> _logger;

    public OutboxDispatcher(AppDbContext context, RabbitMqPublisher publisher, ILogger<OutboxDispatcher> logger)
    {
        _context = context;
        _publisher = publisher;
        _logger = logger;
    }

     public async Task DispatchPendingMessagesAsync()
    {
        var pendingMessages = await _context.OutboxMessages
            .Where(m => !m.Processed)
            .ToListAsync();

        foreach (var message in pendingMessages)
        {
            try
            {
                _logger.LogInformation("Gönderiliyor: {Payload}", message.Payload);
                _publisher.PublishRaw(message.Payload);
                message.Processed = true;
                message.ProcessedAt = DateTime.UtcNow;
                await _context.SaveChangesAsync();
                _logger.LogInformation("Gönderildi: {Payload}", message.Payload);
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Mesaj gönderilemedi: {ex.Message}");
            }
        }
    }
}
