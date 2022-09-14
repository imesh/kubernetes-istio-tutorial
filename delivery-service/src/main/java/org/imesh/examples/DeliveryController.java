package org.imesh.examples;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.sql.Date;
import java.util.ArrayList;
import java.util.List;

@RestController
public class DeliveryController {

    private List<Delivery> deliveries = new ArrayList<>();

    public DeliveryController() {
        deliveries.add(new Delivery("1", "1", Date.valueOf("2022-01-01"), "1st Street, New York, USA", "Delivered"));
        deliveries.add(new Delivery("1", "2",Date.valueOf("2022-01-02"), "2nd Street, Chicago, USA", "Created"));
        deliveries.add(new Delivery("1", "3", Date.valueOf("2022-01-03"), "3rd Street, Sydney, Australia", "Created"));
    }

    @GetMapping("/deliveries")
    public List<Delivery> getDeliveries() {
        return deliveries;
    }

    @GetMapping("/orders/{orderId}/deliveries")
    public List<Delivery> getOrderDeliveries(@PathVariable("orderId") String orderId) {
        List<Delivery> result = new ArrayList<>();
        for (Delivery delivery: deliveries) {
            if(delivery.orderId.equals(orderId)) {
                result.add(delivery);
            }
        }
        return result;
    }
}
