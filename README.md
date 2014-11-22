Nagios check for the Nexus Repos in Go.
=========================================

    Usage of ./check_nexus_repos:
      -c=20: critical level for job queue depth
      -h="http://ci.walmartlabs.com/nexus/service/local": base url for jenkins  like http://ci.walmartlabs.com/nexus/service/local
      -v=false: verbose output
      -w=10: warning level for job queue depth

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

