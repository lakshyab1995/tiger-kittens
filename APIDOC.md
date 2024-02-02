# Tiger Sightings GraphQL API Documentation

## Table of Contents

- [Tiger Sightings GraphQL API Documentation](#tiger-sightings-graphql-api-documentation)
  - [Table of Contents](#table-of-contents)
  - [GraphQL Endpoint](#graphql-endpoint)
  - [APIs](#apis)
      - [Create User](#create-user)
        - [Request-](#request-)
        - [Response-](#response-)
      - [Login](#login)
        - [Request-](#request--1)
        - [Response-](#response--1)
      - [Create Tiger](#create-tiger)
        - [Request-](#request--2)
        - [Response-](#response--2)
      - [List all Tigers](#list-all-tigers)
        - [Request-](#request--3)
        - [Response-](#response--3)
      - [Create Sighting of a Tiger](#create-sighting-of-a-tiger)
        - [Request-](#request--4)
        - [Response-](#response--4)
      - [List Sightings of Tiger](#list-sightings-of-tiger)
        - [Request-](#request--5)

## GraphQL Endpoint

The GraphQL API is accessible at `/` for all queries.
All create endpoints are authenticated, `Authorization` header is required.

## APIs

#### Create User

##### Request-

```
mutation {
  CreateUser(Username: "testuser", Password: "testpassword", Email: "testuser@example.com") {
    ID
    Username
    Email
    Token
    Expiry
  }
}
```

##### Response-

```
{
  "data": {
    "CreateUser": {
      "ID": "8c3ffdc6-4c2d-4761-beac-7d0db531d1b5",
      "Username": "testuser",
      "Email": "testuser@example.com",
      "Token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY5NDg0OTYsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.JBkiHM9OllYA1miTJ2RqvY8mDk-NcPKj4Ix2qQlFC-M",
      "Expiry": "1706948496"
    }
  }
}

```

#### Login

##### Request-

```
mutation {
  Login(input: {Username: "testuser", Password: "testpassword"}) {
    Token
 	Expiry
  }
}

```

##### Response-

```
{
  "data": {
    "Login": {
      "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY5NDg3NTAsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.WiRAHPBBeJ_rvzJ3aTXD5j3u_C09T-JMT-KfD1lbYGk",
      "Expiry": "1706948750"
    }
  }
}

```

#### Create Tiger

##### Request-

```
mutation {
  createTiger(Name: "TigerName", DateOfBirth: "2000-01-01", LastSeenTimestamp: "2024-02-01T00:00:00Z", Coordinates: {Lat: 40.7128, Lon: 74.0060}) {
    ID
    Name
    DateOfBirth
    LastSeenTimestamp
    LastSeenCoordinates {
      Lat
      Lon
    }
  }
}

```

##### Response-

```
{
  "data": {
    "createTiger": {
      "ID": 3,
      "Name": "TigerName",
      "DateOfBirth": "2000-01-01",
      "LastSeenTimestamp": "2024-02-01T00:00:00Z",
      "LastSeenCoordinates": {
        "Lat": 40.7128,
        "Lon": 74.006
      }
    }
  }
}

```

#### List all Tigers

##### Request-

```
{
  listTigers(first: 10) {
    edges {
      node {
        ID
        Name
        LastSeenTimestamp
        LastSeenCoordinates{
          Lat
          Lon
        }
      }
      cursor
    }
    pageInfo {
      hasNextPage
      endCursor
    }
  }
}

```

##### Response-

```
{
  "data": {
    "listTigers": {
      "edges": [
        {
          "node": {
            "ID": 3,
            "Name": "TigerName",
            "LastSeenTimestamp": "2024-02-01T05:30:00+05:30",
            "LastSeenCoordinates": {
              "Lat": 40.7128,
              "Lon": 74.006
            }
          },
          "cursor": "Mw=="
        }
      ],
      "pageInfo": {
        "hasNextPage": false,
        "endCursor": null
      }
    }
  }
}

```

#### Create Sighting of a Tiger

##### Request-

```
mutation {
  CreateSighting(
    TigerID: 3,
    Timestamp: "2022-02-01T14:30:00Z",
    Coordinates: {
      Lat: 100.7128,
      Lon: 174.0060
    },
    ImageURL: "http://example.com/image.jpg"
  ) {
    ID
    TigerID
    Timestamp
    Coordinates {
      Lat
      Lon
    }
    ImageURL
  }
}

```

##### Response-

```
{
  "data": {
    "CreateSighting": {
      "ID": 9,
      "TigerID": 3,
      "Timestamp": "2022-02-01T14:30:00Z",
      "Coordinates": {
        "Lat": 100.7128,
        "Lon": 174.006
      },
      "ImageURL": "http://example.com/image.jpg"
    }
  }
}

```

#### List Sightings of Tiger

##### Request-

```
{
  "data": {
    "ListSightings": {
      "edges": [
        {
          "node": {
            "ID": 8,
            "TigerID": 3,
            "Timestamp": "2024-02-01T05:30:00+05:30",
            "Coordinates": {
              "Lat": 40.7128,
              "Lon": 74.006
            },
            "ImageURL": ""
          },
          "cursor": "OA=="
        },
        {
          "node": {
            "ID": 5,
            "TigerID": 3,
            "Timestamp": "2024-02-01T05:30:00+05:30",
            "Coordinates": {
              "Lat": 40.7128,
              "Lon": 74.006
            },
            "ImageURL": ""
          },
          "cursor": "NQ=="
        }
      ],
      "pageInfo": {
        "hasNextPage": true,
        "endCursor": "NQ=="
      }
    }
  }
}

```