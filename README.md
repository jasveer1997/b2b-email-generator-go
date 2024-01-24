# b2b-email-generator-go
 
- Start local server using main.go () 

### List of APIs

1. GET `/domains?size={page_size}&from={start}&search={search_text}` (This API lists registered domains in product in a paginated fashion)
2. GET `/users?size={page_size}&from={start}&search={search_text}` (This API lists registered users in product in a paginated fashion)
3. PUT `/generate_email` (This API generates single user email with name, domain and returns it. This is currently a single user end point)

