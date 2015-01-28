Nagios check for the Nexus Repos in Go.
=========================================

Checks how many repos are in the BLOCKED_AUTO state. This can happen
if something is wrong with the network.

    Usage of ./check_nexus_repos:
      -w=10: warning level for number of blocked repos
      -c=20: critical level for number of blocked repos
      -h="http://gec-maven-nexus.walmart.com/nexus/service/local": base url for nexus
      -v=false: verbose output

It calls the Nexus API, gets the list of repos, and call the API
again for each repo, checking the state of each.

Build it:
---------

  go build

or:

  go build check_nexus_repos.go

### Docker:

Build it in docker for another platform:

docker run -it -v /Users:/Users -w `pwd` google/golang go build check_nexus_repos.go

Nexus API:
------------

The Nexus API is well documented.

https://repository.sonatype.org/nexus-restlet1x-plugin/default/docs/path__repositories_-repositoryId-_status.html

Example:

curl -v -H "Accept: application/json" http://gec-maven-nexus.walmart.com/nexus/service/local/repositories/fusesource/status 

