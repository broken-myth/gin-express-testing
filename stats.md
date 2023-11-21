# Benchmarking Gin (Golang) and NodeJS (Express) servers

``` Environment - OS - Windows, Python requests library for testing```

- *One simple GET request (string response) :*

    - **Gin** : between 6.2ms to 7.5ms
    - **Express** : between 6.9ms to 7.5ms

- *100 GET requests (string response) :*

    - **Gin** : average 325ms
    - **Express** : average 350ms

- *1000 GET requests (string response) :*

    - **Gin** : average 3237ms (over 5 batches of request)
    - **Express** : average 3428ms (over 5 batches of request)

- *One simple POST request (string response (concat of the request strings), request body with two string fields) :*

    - **Gin** : average 7.582ms
    - **Express** : average 8.774ms

    (There could be prone inaccuracies here as the request body is not being parsed in the same way in both the servers.)

- *100 POST requests (string response (concat of the request strings), request body with two string fields) (5 batches) :*

    - **Gin** : average 324.192ms
    - **Express** : average 362.798ms

    (There could be prone inaccuracies here as the request body is not being parsed in the same way in both the servers.)

- *1000 POST requests (string response (concat of the request strings), request body with two string fields) (5 batches) :*

    - **Gin** : average 3563.4ms
    - **Express** : average 4395.83ms

    (There could be prone inaccuracies here as the request body is not being parsed in the same way in both the servers.)

# Final understandings
```
- Gin seems to be faster in all the cases under the environment tested in. 

- Gin could parse the request body faster than Express in post requests or maybe processing the strings is faster in Go than NodeJS that might have provided better results for Go (basically we concatenated the two incoming fields into one string and sent as the response) 
```


