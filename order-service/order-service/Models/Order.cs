namespace order_service.Models;

public class Order
{
    public Guid Id { get; set; } = Guid.NewGuid();
    public string ProductId { get; set; } = string.Empty;
    public int Quantity { get; set; }
    public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
}
