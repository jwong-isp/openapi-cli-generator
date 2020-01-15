# User
hard-coded:
- client ID (from Auth0)
- issuer endpoint (from Auth0, e.g. `https://istreamplanet.auth0.com/`)

user-supplied:
- identifier (e.g. `default`)
- audience (from Auth0, e.g. `https://api.istreamplanet.net`)
- default user identity for `x-istreamplanet-user-identity` header

# Machine-to-machine
hard-coded:
- issuer endpoint (from Auth0, e.g. `https://istreamplanet.auth0.com/`)

user-supplied:
- identifier (e.g. `default`)
- client ID (from Auth0)
- client secret (from Auth0)
- audience (from Auth0, e.g. `https://api.istreamplanet.net`)
- default user identity for `x-istreamplanet-user-identity` header

```sh
$ events-cli auth add-profile <identifier> <audience> <default user identity>

$ events-cli auth add-profile <identifier> <client-id> <client-secret> <audience> <default user identity>

# BREAKING change, CLI version 2.0
$ events-cli auth add-profile <user|m2m> ...
$ events-cli auth add-profile user <identifier> <audience> <identity>
$ events-cli auth add-profile m2m <identifier> <client-id> <client-secret> <audience> <identity>
```

# Questions
1. How do we configure auth during setup?
2. How do we select auth type at runtime?



# Code Modifications should do:
1. Custom parameters per auth type when creating a profile
2. Figure out at runtime which auth type to use
3. Calling the right middleware/hooks/handlers for the selected scheme


# CLI Current Flow
## ...the code...
1. Inits the cli -- this creates the Root cmd. You can now attach sub cmds to it.
2. Inits the authentication mechanism -- this 
    1. adds the `auth add` cmds, sets up writing to a file
    2. hardcodes the oauth flow to m2m or pkce   *We want to change this.
    3. Adds hooks to add the Authorization header to each request.
    
3. Register the generated code(ucs, events, audit, etc...).

## ...the user...
4. Calls `events-cli auth add ...` with args defined in `step 2 above`.
## ...the code...
5. Creates a `profile` and stores it in `.events/credentials.json`
## ...the user...
6. Calls a cmd like `list-slates`
## ...the code...
7. Checks for a token under `.events/cache.json`
8. Uses token or uses info in `.events/credentials.json` to retrieve a token.
9. Executes the command and returns a response to the user
## ...the user...
10. is happy