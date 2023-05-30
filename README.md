# FizzBuzz
Technical Test submission for Leboncoin

## Running project
Clone it (obviously)
Create an `.env` file if you want to specify either the ENV ou SERVER_HTTP_PORT.

>go run main.go

_(You could run it off of Docker, Kubernetes..._)

... You're all set.

## Routes
3 routes here:

- `/health`: check server status (running -> http code OK 200, not running -> 404)
- `/fizzbuzz`:
  - Accepts five parameters: three integers **int1**, **int2** and **limit**, and two strings **str1** and **str2**.
    - Usage example: `fizzbuzz?int1=3&int2=5&lim=100&str1=fizz&str2=buzz`
    - Reminder : _lim > str1, str2 and all of them must be greater than 0_
- [WIP] `/stats`: 
  - Return the parameters corresponding to the most used request, as well as the number of hits for this request"
