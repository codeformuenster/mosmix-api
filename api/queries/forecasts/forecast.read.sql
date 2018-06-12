SELECT ST_X(p.the_geom) AS lon, ST_Y(p.the_geom) AS lat, ST_Z(p.the_geom) AS alt, f.* FROM forecast_places AS p JOIN forecasts_all AS f ON p.id = f.place_id WHERE id = '{{.id}}'
