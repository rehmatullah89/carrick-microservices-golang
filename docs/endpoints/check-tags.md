# Get tag

**URL** : `/check-tags/{publisher_hash}
```
publisher_hash: publisher unique identifier - required
```
**Example**:
```
GET /check-tags/c4a675af07781d21b0488cb1fc23cea11
```

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "status": true,
    "data": 100,
    "error": null
}
```

## Error Response

**Condition** : If publisher not found

**Code** : `404 NOT FOUND`

**Content** :

```json
{
    "status": false,
    "data": null,
    "error": "Publisher not found by hash: [publisher_hash]"
}
```