FORMAT: 1A
HOST: http://localhost:8080/

# Server 1

Auth backend server

## REGISTER [/register]

### Create New User [POST]

add new user to registry if username not exist

+ Request (application/json)

        {
            "user_name":"username1",
            "password":"password",
            "name":"angling z",
            "role":"owner"
        }
        

+ Response 200 (application/json)

    + Body

            {
                "message": "register succes, now you can login.",
                "username": "username1",
            }

## Authentication [/login]

### Login [POST]

authenticate user

+ Request (application/json)

        {
            "user_name":"username1",
            "password":"password",
        }
        

+ Response 200 (application/json)

    + Body

            {
               "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ1c2VybmFtZTEiLCJuYW1lIjoiYW5nbGluZyB6Iiwicm9sZSI6Im93bmVyIiwidGltZXN0YW1wIjoiMTU4MjU0MDczNSIsImV4cCI6MTYxNDA3NzM3NH0.SNyTFamGc_82QZ7D3gky-pmKqXOwdGxEd4SBYIpCqzQ"
            }



## JWT [/token/check]

### CLAIMS [GET]

Accept jwt token and return private claims if token is valid 

+ Request

    + Headers

            Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ1c2VybmFtZTEiLCJuYW1lIjoiYW5nbGluZyB6Iiwicm9sZSI6Im93bmVyIiwidGltZXN0YW1wIjoiMTU4MjU0MDczNSIsImV4cCI6MTYxNDA3NzM3NH0.SNyTFamGc_82QZ7D3gky-pmKqXOwdGxEd4SBYIpCqzQ
        
+ Response 200 (application/json)

        {
            "private claims": {
                "exp": 1614077374,
                "name": "angling z",
                "role": "owner",
                "timestamp": "1582540735",
                "user_name": "username1"
            }
        }
