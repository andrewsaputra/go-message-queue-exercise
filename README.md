by [andrewsaputra](https://github.com/andrewsaputra)

# Message Queue System

## Summary

Design and implement a simple  message queueing system using Go programming language and RabbitMQ or Kafka.

The system includes the following parts :

#### API
Design an API where it should receive a product data and store it in the database, below are the parameters that should be passed in the API

- user_id
- product_name
- product_description (text)
- product_images (array of image urls)
- product_price (number)

#### Producer
After storing the product details in the database, `product_id` should be passed on the message queue.


#### Consumer
Based on the `product_id`, images should be download and compressed and stored in local. After storing, a local location path should be added as an array value in the `products` table in the `compressed_images` column.

## Database Schema

#### Users

| Column | Type | Notes |
| --- | --- | --- |
| id | INT | PRIMARY KEY |
| name | VARCHAR(255) | |
| email | VARCHAR(255) | |
| created_at | BIGINT | |
| updated_at | BIGINT | |


#### Products
(data should be added from API only)

| Column | Type | Notes |
| --- | --- | --- |
| id | INT | PRIMARY KEY |
| user_id | INT | FOREIGN KEY |
| name | VARCHAR(255) | |
| description | TEXT | |
| images | TEXT[] | |
| compressed_images | TEXT[] | |
| price | NUMERIC(10,2) | |
| created_at | BIGINT | |
| updated_at | BIGINT | |


## Testing
Include unit tests, integration and benchmark tests.
