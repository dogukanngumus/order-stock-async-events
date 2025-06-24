using Microsoft.EntityFrameworkCore;
using order_service.Messaging;
using order_service.Models;
using order_service.Outbox;
using order_service.Services;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("Postgres")));

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddSingleton<RabbitMqPublisher>();
builder.Services.AddScoped<OrderService>();
builder.Services.AddHostedService<order_service.BackgroundServices.OutboxDispatcherBackgroundService>();
builder.Services.AddScoped<OutboxDispatcher>();

var app = builder.Build();

// Polly kullanılabilirdi
using (var scope = app.Services.CreateScope())
{
    var db = scope.ServiceProvider.GetRequiredService<AppDbContext>();

    var maxRetries = 10;
    var retryDelay = TimeSpan.FromSeconds(5);
    var retries = 0;

    while (true)
    {
        try
        {
            db.Database.Migrate();
            SeedData.SeedProducts(app);
            break;
        }
        catch (Exception ex)
        {
            retries++;
            if (retries >= maxRetries)
            {
                Console.WriteLine($"❌ Veritabanına bağlanılamadı: {ex.Message}");
                throw;
            }

            Console.WriteLine($"🔁 Veritabanına bağlanılamadı ({retries}/{maxRetries}): {ex.Message}");
            Task.Delay(retryDelay).Wait();
        }
    }
}

app.UseSwagger();
app.UseSwaggerUI();
app.MapControllers();

app.Run();
