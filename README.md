# Usage

Check https://mosmix-api.codeformuenster.org

# Deployment

Database, api and processor is running inside containers. You need Docker, docker-compose and systemd (for scheduling). The supplied `docker-compose.yml` brings everything needed for running and serving the whole stack

## Steps to recreate the deployment

- Install docker, docker-compose
- Copy the files from the `deployment` folder to you server
  - folder `mosmix` to `/etc/mosmix`
  - folder `systemd/system` to `/etc/systemd/system`
- Start the database on the server
  - `docker-compose -f /etc/mosmix/docker-compose.yml up -d mosmix-postgis`
- Execute the processor(s) once
  - `docker-compose -f /etc/mosmix/docker-compose.yml run --rm processor mosmix_s`
  - `docker-compose -f /etc/mosmix/docker-compose.yml run --rm processor mosmix_l`
- Enable and start the systemd timer on the server
  - `systemctl daemon-reload`
  - `systemctl start mosmix-processor-l.timer mosmix-processor-s.timer`
  - `systemctl enable mosmix-processor-l.timer mosmix-processor-s.timer`
- Start the frontend-proxy and api on the server
  - `docker-compose -f /etc/mosmix/docker-compose.yml up -d api frontend-proxy`

You now should have a running version of the api accessible at port 80.
