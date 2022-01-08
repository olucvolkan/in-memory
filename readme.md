# Golang Test Case

## How to access the app?

```
curl --location --request POST 'https://vast-beyond-51726.herokuapp.com/in-memory/' \
--header 'Content-Type: application/json' \
--data-raw '{
        "key": "testKey",
        "value": "testValue"
}' | jq
```

```
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100    96  100    38  100    58     56     86 --:--:-- --:--:-- --:--:--   149
{
"key": "testKey",
"value": "testValue"
}
```

```
curl --location --request GET 'https://vast-beyond-51726.herokuapp.com/in-memory?key=testkey' | jq
```

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0     55      0 --:--:-- --:--:-- --:--:--    57
{
  "key": "testkey",
  "value": "testValue"
}
```

## How to run locally

```
 ./start.sh
```

## How to test locally

```
 go test -v
```