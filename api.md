# API Documentation

## Resources
/health

/stations
/stations/{id}

/stations/{id}/modules
/stations/{id}/modules/{id}

/stations/{id}/research
/stations/{id}/research/labs
/stations/{id}/research/labs/{id}
/stations/{id}/research/projects
/stations/{id}/research/projects/{id}

## URIs
`/health`
* Endpoint server system health and metadata

`/stations`
* GET: Retrieve list of player-owned stations
* POST: Create a new station; returns station ID

`/stations/{id}`
* GET: Retrieve details about a given station, such as resource metrics

`/stations/{id}/modules`
* GET: Retrieve list of IDs for all modules attached to the station
* POST: Build a new module onto the station

`/stations/{id}/modules/{id}`
* GET: List all details for the given module (e.g. occupants, production)
* PATCH: Update the given module's configuration 
* DELETE: Scrap the module, salvaging materials in the process

`/stations/{id}/research`
* GET: Retreive tech summary

`/stations/{id}/research/labs`
* GET: Retrieve list of research labs and current projects

`/stations/{id}/research/labs/{id}`
* GET: Retrieve details about a given lab and its progress toward a project
* PATCH: Set the current research project for the given lab

`/stations/{id}/research/projects`
* GET: Retrieve full list of available research projects

`/stations/{id}/research/projects/{id}`
* GET: Retrieve details about a given project, including name, cost, and
  benefit
