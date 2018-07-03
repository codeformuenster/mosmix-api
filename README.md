# Deployment

Database, api and processor is running inside containers. You need Docker, docker-compose and systemd (for scheduling). The supplied `docker-compose.yml` assumes an external frontend proxy.

## Steps to recreate the deployment

- Install docker, docker-compose
- Copy the files from the `deployment` folder to you server
  - `docker-compose.yml` to `/srv/mosmix`
  - `mosmix-processor.service` to `/etc/systemd/system/mosmix-processor.service`
  - `mosmix-processor.timer` to `/etc/systemd/system/mosmix-processor.timer`
- Start the database and api on the server
  - Go to `/srv/mosmix`
  - `docker-compose up -d mosmix-postgis api`
- Execute the processor once
  - `docker-compose run --rm processor`
- Enable and start the systemd timer on the server
  - `systemctl daemon-reload`
  - `systemctl start mosmix-processor.timer`
  - `systemctl enable mosmix-processor.timer`

You now should have a running version of the api accessible on port 3000 (or through the url configured in your front end proxy).
