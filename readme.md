## Install and run

- 'git clone https://gitlab.com/AdVonCommerce/carrick-bend/docker-services and follow instruction in readme.md'

### Run using docker-compose
- `docker-compose up --build`

### Add function to database (postgres)
- ```
  CREATE OR REPLACE FUNCTION traffic_source_by_domain(in_domain varchar) RETURNS TABLE(id bigint, name varchar, clid varchar, is_amp bool, is_default bool) LANGUAGE plpgsql AS $$
  DECLARE
      traffic_source_domain_row RECORD;
  BEGIN
      FOR traffic_source_domain_row IN
          select domain, traffic_source_id from traffic_source_domains
      LOOP
          IF in_domain SIMILAR TO '(%.' || traffic_source_domain_row.domain || '|' || traffic_source_domain_row.domain || ')' THEN
              RETURN QUERY select ts.id, ts.name, ts.clid, ts.is_amp, ts.is_default from traffic_sources as ts where ts.id = traffic_source_domain_row.traffic_source_id;
              RETURN;
          END IF;
      END LOOP;
  
      RETURN QUERY select ts.id, ts.name, ts.clid, ts.is_amp, ts.is_default from traffic_sources as ts where ts.is_default = true;
  END;
  $$
  ;```
  
## Endpoints
- [Get tag](docs/endpoints/get-tag.md) : `GET /get-tag/{publisher_hash}`
- [Send tracking](docs/endpoints/send-tracking.md) : `POST /send-tracking/{publisher_hash}`
- [Save visit](docs/endpoints/visit.md) : `POST /visit/{publisher_hash}`
- [Get count of unused tags](docs/endpoints/check-tags.md) : `GET /check-tags/{publisher_hash}`