SELECT id, name, ST_X(the_geom) AS lon, ST_Y(the_geom) AS lat, ST_Z(the_geom) AS alt FROM forecast_places WHERE name ILIKE '%{{.name}}%'
