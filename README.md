oaizz
=====

Command line OAI-harvesting. The goal is to enable access to
data from various repositories and time ranges.

Data is cached locally and most of the time transparent to the user.

The program can be used with cron to keep an up-to-date version
of the havested data available.

----

Show info about OAI endpoint.

    $ oaizz -v -url www.doabooks.org/oai

Download all available data.

    $ oaizz -url www.doabooks.org/oai

Download a collection only.

    $ oaizz -url www.doabooks.org/oai -set articles

Download with a given prefix.

    $ oaizz -url www.doabooks.org/oai -set articles -prefix marcxml

Download for a given date range.

    $ oaizz -url www.doabooks.org/oai -set articles -prefix marcxml -from 2010-01-01 -to 2011-01-01

Emit the file path for a given harvest.

    $ oaizz -files -url www.doabooks.org/oai -set articles -prefix marcxml -from 2010-01-01 -to 2011-01-01


This program is a tool for the service provider side of the OAI protocol (section 2.1).

----

OAI in short
------------

Three entities: resources (real-world), items (abstract), records (formatted). 

Items have an identifier, which is unique per repository. There can be multiple formats. The ID must be a URI, but scheme can only be used when adhering to standards.

ID plays a role in:

* ListIdentifier, ListRecords
* GetRecord

The response is always XML.

Response has a header, metadata and about.

