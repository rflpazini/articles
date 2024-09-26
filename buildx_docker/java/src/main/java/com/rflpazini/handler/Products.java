package com.rflpazini.handler;

import io.micronaut.http.annotation.Controller;
import io.micronaut.http.annotation.Get;
import io.micronaut.serde.annotation.Serdeable;

import java.util.List;

@Controller("/products")
class Products {

    @Get("/")
    public List<Product> listProducts() {
        return List.of(
                new Product(1, "Laptop", 999.99),
                new Product(2, "Smartphone", 699.99)
        );
    }
}

@Serdeable
record Product(int id, String name, double price) {}