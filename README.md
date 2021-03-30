***
# NOTICE:

## This repository has been archived and is not supported.

[![No Maintenance Intended](http://unmaintained.tech/badge.svg)](http://unmaintained.tech/)
***
NOTICE: SUPPORT FOR THIS PROJECT HAS ENDED 

This projected was owned and maintained by Walmart. This project has reached its end of life and Walmart no longer supports this project.

We will no longer be monitoring the issues for this project or reviewing pull requests. You are free to continue using this project under the license terms or forks of this project at your own risk. This project is no longer subject to Walmart's bug bounty program or other security monitoring.


## Actions you can take

We recommend you take the following action:

  * Review any configuration files used for build automation and make appropriate updates to remove or replace this project
  * Notify other members of your team and/or organization of this change
  * Notify your security team to help you evaluate alternative options

## Forking and transition of ownership

For [security reasons](https://www.theregister.co.uk/2018/11/26/npm_repo_bitcoin_stealer/), Walmart does not transfer the ownership of our primary repos on Github or other platforms to other individuals/organizations. Further, we do not transfer ownership of packages for public package management systems.

If you would like to fork this package and continue development, you should choose a new name for the project and create your own packages, build automation, etc.

Please review the licensing terms of this project, which continue to be in effect even after decommission.

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

