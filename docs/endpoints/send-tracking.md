# Save tracking info

...

**URL** : `/send-tracking/{publisher_hash}`
```
publisher_hash="[publisher unique identifier]"
```
**Example**:
```
/send-tracking/c4a675af07781d21b0488cb1fc23cea11
```

**Method** : `POST`

**Data constraints**

Provide name of Account to be created.

```json
{
    "t": "[tag]", // required
    "ci_t": "[click identifier type]", // required if a_u is empty
    "ci_v": "[click identifier value]", // required if ci_t is not empty
    "r": "[referrer]",
    "u_agent": "[user agent]",
    "c_u": "[click url]", // required
    "a_u": "[current page url]", // required if ci_t is empty
    "amp": "[is amp page]"
}
```

**Data example**

```json
{
    "t": "tag123",
    "ci_t": "gclid",
    "ci_v": "gclid123",
    "r": "http://google.com",
    "u_agent": "USER_AGENT",
    "c_u": "http://amazon.com/asin",
    "a_u": "http://store.com",
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

**Condition** : If history item not updated

**Code** : `500 Update error`

**Content** : 
```json
{
    "status": false,
    "data": null,
    "error": "Update error"
}
```

### Or

**Condition** : If history item not created

**Code** : `500 Create error`

**Content** : 
```json
{
    "status": false,
    "data": null,
    "error": "Create error"
}
```
