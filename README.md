# Linkme

Simple link management platform.

we just run 
for seeder examples runs a one-time database population script
go run cmd/seed/main.go

for main = runs your API
go run cmd/server/main.go


curl commands

create a new user 
POST
Register 

curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"email":"newuser@example.com","password":"password123"}'

login 

curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"email":"newuser@example.com","password":"password123"}'

link - to see links we entered for that 

curl -X GET http://localhost:8080/links \
-H "Authorization: Bearer YOUR_TOKEN_HERE"

Create a link 

curl -X POST http://localhost:8080/links \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_TOKEN" \
-d '{"title":"OpenAI","url":"https://openai.com"}'

Delete 

curl -X DELETE http://localhost:8080/links/1 \
-H "Authorization: Bearer YOUR_TOKEN"

Put or update 

curl -X PUT http://localhost:8080/links/1 \
-H "Authorization: Bearer YOUR_TOKEN" \
-H "Content-Type: application/json" \
-d '{"title":"OpenAI Updated","url":"https://openai.com"}'


#######################