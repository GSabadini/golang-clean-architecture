db = db.getSiblingDB('challenge');

db.createUser({
    user: 'dev',
    pwd: 'dev',
    roles: [
        {
            role: 'root',
            db: 'admin',
        },
    ],
});

users = db.createCollection('users');
// db.users.createIndex( { "document.type": 1 }, { unique: true } )
// db.users.createIndex( { "id": 1 }, { unique: true } )
// db.users.createIndex( { "email": 1 }, { unique: true } )


createCollection('transfers');