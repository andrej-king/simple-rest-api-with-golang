#!/bin/bash
mongo <<EOF
use $MONGO_API_DB_NAME
db.createUser({
  user: '$MONGO_API_DB_USERNAME',
  pwd: '$MONGO_API_DB_PASSWORD',
  roles: [{
    role: 'readWrite',
    db: '$MONGO_API_DB_NAME'
  }]
})
db.createCollection('$MONGO_API_DB_COLLECTION')
EOF
