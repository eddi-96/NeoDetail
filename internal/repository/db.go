package repository

import "database/sql"

// ÚNICA definición global de DB para todo el package repository
var DB *sql.DB
