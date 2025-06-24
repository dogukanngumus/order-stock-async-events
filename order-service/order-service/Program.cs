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
using (var scope = app.Services.CreateScope())
{
    var db = scope.ServiceProvider.GetRequiredService<AppDbContext>();
    db.Database.Migrate();
    SeedData.SeedProducts(app);
}

app.UseSwagger();
app.UseSwaggerUI();
app.MapControllers();
app.Run();