# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

### Mirroring Product APIs
There are several products (Gloo Edge, Gloo Platform) which leverage this repository to host read-only mirrors of their APIs.

To keep those APIs distinct for each product and each long term support version of that product, we use github branches to maintain isolation.

The branch naming pattern is: [PRODUCT PREFIX]-[PRODUCT LTS BRANCH NAME]

So for example, the Gloo Edge 1.14.x APIs, are maintained in the `gloo-v1.14.x` branch in this repo. 


### Main Branch
As a result of the above API tracking approach, the `main` branch in this repository has no functionality.
