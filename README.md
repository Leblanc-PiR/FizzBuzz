# 🍾 FizzBuzz
Technical Test submission for Leboncoin

## 🏃 Running project
Clone it (obviously)
Create an `.env` from `.env.example` file if you want to specify either the ENV or SERVER_HTTP_PORT...

[Compile and install](https://go.dev/doc/tutorial/compile-install)

```
> go build
> ./FizzBuzz
```

Or, if you want to do it locally
```
> go run main.go
```

Workflow was integrated using builtIn GitHub tools.

_(You could run it off of Docker, Kubernetes...)_

... You're all set, the server is running on specified port !

## 🛣️ Routes
3 routes here:

- `/health`: check server status (running -> http code OK 200, not running -> 404)
- `/fizzbuzz`:
  - Accepts five parameters: three integers **int1**, **int2** and **limit**, and two strings **str1** and **str2**.
    - Usage example: `fizzbuzz?int1=3&int2=5&lim=100&str1=fizz&str2=buzz`
    - Reminder : _lim > str1, str2 and all of them must be greater than 0_
- `/stats`: 
  - Return the parameters corresponding to the most used request(s), as well as the number of hits for this request"

## 📈 Ways to improve
- Did not go for TDD here, probably should have.
- Makefile, Dockerfile or else
- Connect to a real DB, use Squirrel
- Front could be fancier 🙃
