# Get tag

**URL** : `/get-tag/{publisher_hash}?a_u=[current page url]&r=[referrer]&t_type=[traffic_type]`
```
publisher_hash: publisher unique identifier - required
a_u: current page url - required for organic traffic
r: http referrer - required for orgranic traffic
t_type: traffic type - required (organic or paid)
```
**Example**:
```
Paid traffic: /get-tag/c4a675af07781d21b0488cb1fc23cea11/?t_type=paid
Organic traffic: /get-tag/c4a675af07781d21b0488cb1fc23cea11?a_u=https://publisher.com/review1&r=https://google.com&t_type=organic
```

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "status": true,
    "data": "tag-example",
    "error": null
}
```

## Error Response

**Condition** : If publisher not found or tag not found

**Code** : `404 BAD REQUEST`

**Content** :

```json
{
    "status": false,
    "data": null,
    "error": "Tag not found"
}
```