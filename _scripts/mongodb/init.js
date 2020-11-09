db = db.getSiblingDB('challenge');

users = db.createCollection('users');
db.users.createIndex( { "document.type": 1 }, { unique: true } )
db.users.createIndex( { "email": 1 }, { unique: true } )

