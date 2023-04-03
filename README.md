# ASSIST Nice Things Backend

## Dayton Berezoski

This project is the backend for the overarching nice things website. It utilizes the Gin framework along with GORM to create a REST API that provides all the services that will be needed.

## Routes

```
/api
├── /users
│   ├── /signUp
│   ├── /signIn
│   ├── /signOut
│   ├── /changePassword
│   └── /getUsers
└── /niceThings
    ├── /createNiceThing
    ├── /editNiceThing
    ├── /deleteNiceThing
    └── /generateNiceThing
```
