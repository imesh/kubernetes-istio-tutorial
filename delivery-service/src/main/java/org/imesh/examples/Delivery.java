package org.imesh.examples;

import java.util.Date;

public class Delivery {
    public String id;
    public String orderId;
    public Date date;
    public String address;
    public String status;

    public Delivery(String id, String orderId, Date date, String address, String status) {
        this.id = id;
        this.orderId = orderId;
        this.date = date;
        this.address = address;
        this.status = status;
    }
}
