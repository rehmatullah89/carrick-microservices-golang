# Save visit info

...

**URL** : `/visit/{publisher_hash}`
```
publisher_hash="[publisher unique identifier]"
```
**Example**:
```
/visit/c4a675af07781d21b0488cb1fc23cea11
```

**Method** : `POST`

**Data constraints**

Provide name of Account to be created.

```json
{
    "t": "[tag]",
    "ci_t": "[Click identifier type]", // required if a_u is empty
    "ci_v": "[Click identifier value]", // required if ci_t not empty
    "r": "[referrer]",
    "a_u": "[current page url]", // required if ci_t is empty
    "u_agent": "[user agent]", // required
    "amp": "[is amp page]"
}
```

**Data example**

```json
{
    "ci_t": "gclid",
    "ci_v": "gclid123",
    "r": "https://google.com",
    "a_u": "https://domain.com",
    "u_agent": "Mosila",
    "t": "tag-1",
    "amp": "1"
}
```

## Success Response

**Condition** : If everything is OK

**Code** : `200 OK`

**Content example**

```json
{
    "status": true,
    "data": null,
    "error": null
}
```

## Error Responses

**Condition** : If publisher not found

**Code** : `404 Not Found`

**Content** : 
```json
{
    "status": false,
    "data": null,
    "error": "Publisher not found"
}
```

### Or

**Condition** : If request validation error

**Code** : `400 Bad Request`

**Content** : 
```json
{
    "status": false,
    "data": null,
    "error": "Bad Request"
}
```

### Or

**Condition** : If visit item not created

**Code** : `500 Create error`

**Content** : 
```json
{
    "status": false,
    "data": null,
    "error": "Create error"
}
```
