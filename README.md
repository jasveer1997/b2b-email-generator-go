# b2b-email-generator-go

- Deployment is carried through vercel cloud - it restricts code for some basic impl - like routes are to be defined in api/index, a vercel.json for config. 
- Start local server using main.go (IMP: This file is useless during deployment) 

### List of API

1. GET `/domains?size={page_size}&from={start}&search={search_text}` (This API lists registered domains in product in a paginated fashion)
2. GET `/users?size={page_size}&from={start}&search={search_text}` (This API lists registered users in product in a paginated fashion)
3. ~~PUT~~ POST `/generate_email` (This API generates single user email with name, domain and returns it. This is currently a single user end point. POST used instead of PUT due to vercel library limitation)

