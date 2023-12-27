# Go JSON Web Token

##  Run Redis if needed

If you want to use Redis with a token blacklist

```bash
docker-compose up
```

## Run main

```bash
go mod tidy
go run main.go
```


## Curl command

### Signup

![Signup](./docs/signup.png)

#### Request
```bash
curl --location --request POST '0.0.0.0:8080/api/v1/signup' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "name": "jun",
    "email": "test@alamoa.io",
    "password": "password"
}'
```
#### Response

```bash
{
    "Message": "signup success"
}
```

### Login

![Login](./docs/login.png)

#### Request

```bash
curl --location --request POST '0.0.0.0:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "jun",
    "email": "test@alamoa.io",
    "password": "password"
}'
```

#### Response

```bash
{
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNzA4MDQ1LCJpYXQiOjE3MDM3MDc5ODUsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.Qdn9o8evufrjkb4izNhBJEGQyWIrnJ0rCbIOlgK-yp4FVpHavgoZOEhWTYt64vig9yTALZ6PF_vvwmnUk-TH68OaBQQuPXFK2g4aGIauf1-cR2_QKbf3ut8DvFD0ZFQfClCv84Y1cBvtAGXJd8BoqBVUjE3cNPD9GFRZUdM7ElSK_8nlVVc0VzQO3-17RbzBk2SWGtvAvLiSCuHL3Xbhvqwo6AVsFNw4LfqaKQiaa26coeKSnEYslMEIDnFt4etryoOpPsAc2jwu3KjwxfR8U3CGxRqt-2N1F4_JcIkdoY17ROaRxZx21JibFRL2-Ch1MK7gxvZ-7UlWBhuk8VuHow",
    "refreshToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNzExNTg1LCJqdGkiOiJlMzg1MGJkMi1hYzhkLTRjOTItODY0MS05NDgxOTc0MGQ4NTciLCJpYXQiOjE3MDM3MDc5ODUsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.G78kV1gk79fB4eAWNBRRMvu8o6iHdoP1V3RtUWg1pes3HSObrK1hrIdxrzMui8knN7i54w0a3cHhdHoW9zdFdu3QERKdmlJUsYmPPbitrfNAZs2dh--oOrYdrDIRA6d0nrnXcbPwRDUk5AoP2or7aw-ZwkwLEG-tJaEXTMomVBYYNTkqkLljlHM7PoOjs85jdFaXBmsnZGqnMpKtbBPXnWv3uiuISSSxjSH8rZH20iuhMFyoKUlw4qdayg4ejahOXV25sKgd5UfyFiapIbC8uf5lyA9iMGVRcqdeFzWO1bg0sADAjwg_gRt9nWHUggatkMaIjoRK8FW0qdO8-ROxgQ"
}
```

### Profile

![Profile](./docs/profile.png)

#### Request
```bash
curl --location --request GET '0.0.0.0:8080/api/v1/profile' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNzA4MTI3LCJpYXQiOjE3MDM3MDgwNjcsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.eudAhqCDaZWyyp8C1P-mQNlUqF5G4P7xe_vNCrKg8175kse51TTrDCzoBMpkrBrCIvU09vpHoXZZbDpkKlId_U9dNjfPQuDaK1fkrrBvvqENjYe_a9KdIhVXBaStr4_CGryTraYgKMqcE-xaDCYW0WvbAV6sdu4AS4Nx-D_Wl8au7XMfAyfSDsm8szwHMmgDu2L1NBSSbU8rzM_vVRoPTsRnG7w9q44h-eBw03KwK1hKg9qcfIKoUM334cOqn7YM1wE2z0sd-42CoLPbrARyrzbIIK-p-aap2tKrfLF2tGw6Xim9J7k41uSveulILBmCnhqUX7xSTQ6Qo9sp9NV4-g'``
```

#### Response

```bash
{
    "email": "test@alamoa.io",
    "name": "jun"
}
```

###  Refresh Token

![Refresh](./docs/refresh.png)

#### Request

```bash
curl --location --request POST '0.0.0.0:8080/api/v1/refresh_token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "refreshToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNzExODMzLCJqdGkiOiJmNmJiMjM2ZS00NWFkLTQ0MDgtOWE4Zi1mNDU4ZDQ0MDcxZDgiLCJpYXQiOjE3MDM3MDgyMzMsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.HzjoOoe7FJrW6QLUvQHEIlXxSrYnGcCExTmxA_MyXMrQt8e8jCKN8lTEMERg037XMql3rwoly4arf2DVHroFGE1A8W8Dmq4Z8_1nWVhlxFzKo5zT8verCVtoyJ5PXOPiqjkToRZQXARwUOeSqkZwM_3fX587lLMVOK6PXDGskG4vxPStYOnl3-wPZqybPAyAQcNzfgkpSOMc9d_GSrviZok0iRPD8h7vvERtRfyl5mP9qb20cOS8Nog8GWdGEs_Pe7skuh5z3992MSjd_UEkHIWWlz1uGApDMsHqDrG-JZfb0mQiAUJ8Wny4-FbpkylFuLOLGNDyrwR24ZkL13vCrA"
}'
```

#### Response

```bash
{
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNzA4MzEyLCJpYXQiOjE3MDM3MDgyNTIsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.J7UOuAeCmr-ySAlw8jGzYV2-_WIWuQLR0bgkw1S9Qxec16qtgGzGWA5hY-f_qPAtbikvvKfieBTWn3g2bNb5AlnC9o7Sh4gDyBkJ_TYU_4LblNReb_g61HDFg2OYWWaWVcmuo-l-KLaVWEVWJLvZt-bF_RflCTyA24YiNvJ4mu0N_a_HRvz9ILju8tFoabzqsn7arVnCgzYIFBJT_gXg21BYPgXiaToea5wLY6rfuInPVkYmqLrlSKGaucCQqvgsyaJgbcOXkGwscBq_5rGDcYtPvOwAg253tD6wTFdecKJe37aJ-O72vs03zpBONkCPWW59nRUvWnTajqXiAYlBBA",
    "refreshToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNzExODMzLCJqdGkiOiI2ZjhhOWM2OC1mODIzLTQxNzMtODQwMi05NmIxMWE1MGEzYjIiLCJpYXQiOjE3MDM3MDgyNTIsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.ZPlVOgr-GTC9VW3E62zoRPrlra_F4zbvoiqF-2WuHaboirTU3Q-EX9u5ODYltY9Q_09EYlNZByLqBydQPBQ4XOChIXT_zjKjEq8wYL2To3PnAUOAtjfugsl2A0NtI66DwM0jkF21Bnplv1zkCkPFVlMjgU5yZFZqoEpP94kr_yuMJQhyhvF3enXzVeC8ev4TmCDebuuF5r8U5kEpKr2CabWFqf07aXlrH8va2JbNI3FLuZNgYcZhbdPpy0ZcuL6jtxJswytzFD8oMIsYFLoFozl9u0wsWUs9MTsd21w2dgcv6HeRi8YFfpl50HUUOTS7gwtWz5KqSmqL1td4nKyGQw"
}
```

###  Logout

![Logout](./docs/logout.png)

#### Request

```bash
curl --location --request POST '0.0.0.0:8080/api/v1/logout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InNha2FpanVuc29jY2VyQGdtYWlsLmNvbSIsImV4cCI6MTcwMzQ4ODkyNSwiaXNzIjoiQXV0aFNlcnZpY2UifQ.czZGRRK1-EO9_Gy65FdMkAXTPwJpcQw-uRjwJoFr_Gc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNjQzNjE0LCJpYXQiOjE3MDM2NDM1NTQsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.eGafi413sjONz_RhXaUsHlGV8-x1e32tw6QYYj4rM_cS-DEnaLu0P1A13Bf2eH5VCdkGIObZpPXauWts4C2NWqnOGtTTxj4dGdkOnoRB_a6rKqvODk0hNq-blZvpdJAgANRmprRvbYziZvYXmYTEu9yH904o_xW6YvBEboPBAhX244eLrUREaCxzXdjwKtQ2oXhH5wUcRVe6Z6wpo2EBZXe1JVlY9JGU6aO3P8tVYQsL9ou0dyD6nUAMZ6CS38ltVCJbAORhNR8xZKAqG_7uy75xOX24tuZgR6CzdnNu1r6qTezvghfrChmeaCmLt4jVR3vx9ZNLy6tc9EdAIzyJBA",
    "refreshToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYWxhbW9hLmlvIiwiYXVkIjoiaHR0cHM6Ly9hcGkuYWxhbW9hLmlvIiwiZXhwIjoxNzAzNjQ3MTU0LCJqdGkiOiI3YTcyNTM1ZS1lZmZkLTRmMzgtYjUxMy05NzZiNTExMTg4YzAiLCJpYXQiOjE3MDM2NDM1NTQsImlzcyI6Imh0dHBzOi8vaWRwLmFsYW1vYS5pbyIsInN1YiI6IkFjY2Vzc1Rva2VuIn0.VVq2tFyp7yml18aCkVknnLX85dhPw1buMNAHLElHxziCSUJ0srkCNj11m4OLxW-f3qzJX6xjwVQ5m6oZUy14aTKKUrjT1ifM9_c6_jkIpSnAA3x17QKLlLzytE2cuQfagliRBffa7njZ-QhjwFuPoPsE2w5AqoQacuTt2wUFRSV2p2tzlklz6vhbUHWRGJjNLyZOSxFaydGw8vPatj860dBRh1mcXOLDMSzE90tEvKL-k1mBeUg-qaQIeVtkvY5nr5RrVL3JrQf9_ztsL1I_sXisj-rr_RetMR8Jg7Q_v7Aj2T8pEw2uCoW5dFImRN8ATj7uJXgcamynfoFuLbsW2w"
}'
```

#### Response

```bash
{
    "Message": "logout success"
}
```

## Postman config

[postman config](docs/json_web_token.postman_collection.json)
