using order_service.Models;

namespace order_service.Services;

public static class SeedData
{
    public static void SeedProducts(WebApplication app)
    {
        using var scope = app.Services.CreateScope();
        var context = scope.ServiceProvider.GetRequiredService<AppDbContext>();

        if (!context.Products.Any())
        {
            context.Products.AddRange(
                new Product { Name = "Kalem", Stock = 100 },
                new Product { Name = "Defter", Stock = 50 }
            );

            context.SaveChanges();
        }
    }
}
