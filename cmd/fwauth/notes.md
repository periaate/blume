New link -> link
Use link -> sess
Use sess -> auth

link
- label
- uses
- expiration
- origin
- session duration

session
- label
- expiration
- origin
- cookie hash

store
{root}/{origin}/{hash}.json

api
fw-auth/{origin}/{cookie}
