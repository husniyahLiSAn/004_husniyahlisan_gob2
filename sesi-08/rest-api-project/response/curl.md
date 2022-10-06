## Create
```
curl --location --request POST 'http://localhost:3000/person' \
--form 'first_name="Husniyah"' \
--form 'last_name="Lisan"'
```

## Update
```
curl --location --request PUT 'http://localhost:3000/person?id=2' \
--form 'first_name="Carmellia"' \
--form 'last_name="Maxwell"'
```

## Get All
```
curl --location --request GET 'http://localhost:3000/persons'
```

## Get By Id
```
curl --location --request GET 'http://localhost:3000/person/1'
```

## Delete By Id
```
curl --location --request DELETE 'http://localhost:3000/person/3'
```