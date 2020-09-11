# Simple Mongo-backed GO server to handle user authentication

## Summary
This app is my attempt at creating a mongodb-backed go server to handle user authentication. This was my first time using golang and mongodb so it was a great learning experience for me. After attempting this challenge, I'd like to further my skills in these technologies as I can see there is so much to learn. I used jwt-token authentication and bcrypt because I've used those same technologies in one of my apps created with Node.js and PostgreSQL. As I was studying through this challenge, I did start to notice some similarities in the code.

The biggest challenge for me was integrating mongodb. I had trouble testing the server and database connection locally with MongoDB. Overall, I feel that this code functions for the most part, but does need more tweaking for full functionality.

I used the following tech:

1. Gin web framework for server.
2. mgo MongoDB driver for Go.
3. jwt-go for authentication method and bcrypt for password hashing.

## API Documentation

### POST /register
Creates an user account with account name, password, and email inputs.

### POST /login
Logs user into account with jwt-token authentication.
