# fwauth
fwauth is a forward auth service.

fwauth is made with the assumption that the following statements are true:
1. You are hosting public facing services.
2. You want to restrict access to your services from the public.
3. User based authentication is undesired or infeasible.
4. At least one of the following requirements is true:
	- Authorization must be fleeting, single time, etc.
	- Authorization must be anonymous or userless.
	- Authorization must be persisting.
	- Authorization must be revokable.

fwauth uses links to generate sessions. Sessions to authenticate arbitrary requests.

## License
fwauth is licensed under GPL 3.0.
