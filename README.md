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

1. id - int, primary key
2. name - name of the users
3. phone - contact number of the user
6. created_at
7. updated_at 

#### Products
(data should be added from API only)

1. id - int, primary key
2. name - string, name of the product
3. description - text, about the product
4. images - array
5. price - number
6. compressed_images - array
7. created_at
8. updated_at

## Testing
Include unit tests, integration and benchmark tests.
