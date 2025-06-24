using System;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System.Threading;
using System.Threading.Tasks;
using order_service.Outbox;

namespace order_service.BackgroundServices;

public class OutboxDispatcherBackgroundService : BackgroundService
{
    private readonly IServiceProvider _serviceProvider;
    private readonly ILogger<OutboxDispatcherBackgroundService> _logger;
    private readonly TimeSpan _interval = TimeSpan.FromSeconds(10);

    public OutboxDispatcherBackgroundService(IServiceProvider serviceProvider, ILogger<OutboxDispatcherBackgroundService> logger)
    {
        _serviceProvider = serviceProvider;
        _logger = logger;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        _logger.LogInformation("OutboxDispatcherBackgroundService başlatıldı.");

        while (!stoppingToken.IsCancellationRequested)
        {
            try
            {
                _logger.LogInformation("Outbox mesajları dispatch ediliyor...");

                using var scope = _serviceProvider.CreateScope();
                var dispatcher = scope.ServiceProvider.GetRequiredService<OutboxDispatcher>();

                await dispatcher.DispatchPendingMessagesAsync();

                _logger.LogInformation("Outbox mesajları başarılı şekilde dispatch edildi.");
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Outbox mesajları dispatch edilirken hata oluştu.");
            }

            await Task.Delay(_interval, stoppingToken);
        }

        _logger.LogInformation("OutboxDispatcherBackgroundService durduruluyor.");
    }
}
