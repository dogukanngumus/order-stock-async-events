using Microsoft.AspNetCore.Mvc;
using order_service.Models;
using order_service.Services;

namespace order_service.Controllers;


[ApiController]
[Route("api/[controller]")]
public class OrderController : ControllerBase
{
    private readonly OrderService _orderService;

    public OrderController(OrderService orderService)
    {
        _orderService = orderService;
    }

    [HttpPost]
    public IActionResult CreateOrder([FromBody] Order order)
    {
        var createdOrder = _orderService.CreateOrder(order);
        return Ok(createdOrder);
    }

    [HttpGet]
    public IActionResult GetOrders()
    {
        return Ok(_orderService.GetAllOrders());
    }
}