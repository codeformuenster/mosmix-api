WITH base_qry AS (SELECT
      id, name,
      ROUND(ST_X(the_geom)::numeric, 2) AS lng,
      ROUND(ST_Y(the_geom)::numeric, 2) AS lat,
      ST_Z(the_geom) AS alt,
      the_geom
    FROM {{ .schema }}.forecast_places
    {{ if isSet "name" }}WHERE name ILIKE '%{{ .name }}%'{{ end }})

SELECT id, name, lng, lat, alt{{ if isSet "lat" }}{{if isSet "lng" }}, ROUND((ST_DistanceSphere(the_geom, ST_GeomFromText('POINT({{ .lng }} {{ .lat }})', 4326)) / 1000)::numeric, 2) AS distance_km{{ end }}{{ end }}
FROM base_qry{{ if isSet "lat" }}{{if isSet "lng" }} WHERE
ST_DistanceSphere(the_geom, ST_GeomFromText('POINT({{ .lng }} {{ .lat }})', 4326)) < ({{ defaultOrValue "distance" "10" }} * 1000)
ORDER BY distance_km{{ end }}{{ end }}
