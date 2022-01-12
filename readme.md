# Golang Test Case

## How to access the app?

```
curl --location --request POST 'https://vast-beyond-51726.herokuapp.com/in-memory' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key" : "test",
    "value" : "testvalue"
}' | jq
```

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    84  100    35  100    49     57     80 --:--:-- --:--:-- --:--:--   140
{
  "key": "test",
  "value": "testValue"
}
```

```
curl --location --request GET 'https://vast-beyond-51726.herokuapp.com/in-memory?key=test' | jq
```

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0     55      0 --:--:-- --:--:-- --:--:--    57
{
  "key": "test",
  "value": "testValue"
}
```

```
curl --location --request DELETE 'https://vast-beyond-51726.herokuapp.com/in-memory' \
--header 'Content-Type: application/json'  | jq
```
```
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100    36  100    36    0     0     56      0 --:--:-- --:--:-- --:--:--    58
{
"code": 0,
"msg": "Removed All Data"
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