# repopull
Service that listens on a socket and updates a specified repo. Written for a school project.

### Description
This program will listen on port 8080 for incoming GET requests. The `repoUpdate()`
function will attempt to update repos in the `/var/www/html` directory.

This program was written quickly to fit a specific use case for a school project.
(and was a great excuse to write some golang) Feel free to fork or file an issue.


### Usage
To use, simply hit the endpoint with a web browser or curl.

`curl -v mydomain.com:8080/?repo=MyRepo&branch=MyBranch`

`http://www.mydomain.com:8080/?repo=MyRepo&branch=MyBranch`

Run from cron with:

`curl -sm 30 mydomain.com:8080/?repo=MyRepo&branch=MyBranch`

#### install.sh
This script will install repopull to `/usr/local/bin` and create a systemd service.

#### upgrade.sh
This script will `git pull`, stop and replace the repopull binary and then start the service.

### Disclaimer
This service exposes your web server directly to the internet on port 8080.
Only use this service if you understand the implications that this brings.
