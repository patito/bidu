FORMAT: 1A
HOST: https://localhost/

# Bidu

Cuidado com o bidu, ele pega no seu piru.


## Sites [/api/sites]

This endpoint creates, remove, search for sites.

### Create Site [POST]

+ Attributes

    + name (string, required) - A unique identifier for the client.
    + slug (string, required) - A unique identifier for the client.
    + physical_addresss (string, required) - A unique identifier for the client.
    + comments (string, required) - A unique identifier for the client.

+ Request 

    + Headers
     
            Accept: application/json

    + Body
    
            { 
                "name": "4DTest",
                "slug": "4dtest",
                "physical_address": "brighton",
                "comments": "comments"
            }
        
+ Response 201

## Fetch a Site [/api/sites/{name}]

This endpoint retrieves the site object, if it exists.

### Fetch Site [GET]

+ Parameters

    + name: 4DTest (string, required) - A unique identifier for the site.

+ Request 

    + Headers
     
            Accept: application/json

+ Response 200 (application/json)

    + Body
    
            {
                "id": "1",
                "name": "4DTest",
                "slug": "4dtest",
                "physical_address": "brighton",
                "comments": "comments"
            }

## Update Site [/api/sites/{name}]

This endpoint removes an specific, if it exists.

### Update Site [PUT]

+ Parameters

    + name: 4DTest (string, required) - A unique identifier for the site.

+ Request 

    + Headers
     
            Accept: application/json

    + Body

            {
                "name": "5DTest",
                "slug": "5dtest",
                "physical_address": "5D brighton",
                "comments": "5D comments"
            }

+ Response 200

## Remove Site [/api/sites/{name}]

This endpoint removes an specific, if it exists.

### Delete Site [DELETE]

+ Parameters

    + name: 5DTest (string, required) - A unique identifier for the site.

+ Request 

    + Headers
     
            Accept: application/json

+ Response 200
