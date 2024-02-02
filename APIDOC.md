# Tiger Sightings GraphQL API Documentation

## Table of Contents

- [Tiger Sightings GraphQL API Documentation](#tiger-sightings-graphql-api-documentation)
  - [Table of Contents](#table-of-contents)
  - [GraphQL Endpoint](#graphql-endpoint)
  - [APIs](#apis)
      - [Create User](#create-user)
      - [Login](#login)
      - [Create Tiger](#create-tiger)
      - [List all Tigers](#list-all-tigers)
      - [Create Sighting of a Tiger](#create-sighting-of-a-tiger)
      - [List Sightings of Tiger](#list-sightings-of-tiger)

## GraphQL Endpoint

The GraphQL API is accessible at `/` for all queries.

## APIs

#### Create User

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

#### Login

```
mutation {
  Login(input: {Username: "testuser", Password: "testpassword"}) {
    Token
 	Expiry
  }
}

```

#### Create Tiger

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

#### List all Tigers

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

#### Create Sighting of a Tiger

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

#### List Sightings of Tiger

```
{
  ListSightings(TigerID: 3, first: 10) {
    edges {
      node {
        ID
        TigerID
        Timestamp
        Coordinates{
          Lat
          Lon
        }
        ImageURL
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