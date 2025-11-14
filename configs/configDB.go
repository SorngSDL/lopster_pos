package configs

const MongoURI = "mongodb+srv://admin:password11@cluster0.bj4mmff.mongodb.net/?appName=Cluster0"
const DBName = "phinix_khung"
const UserCollection = "users"
const CategoryCollection = "category"

var JWTSecret = []byte("super-secret-key-change-this")

const AccessTokenTTL = 30
const RefreshTokenTTL = 7
