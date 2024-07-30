from flask import Flask, jsonify, request
from flask_pymongo import PyMongo
from flask_jwt_extended import JWTManager, create_access_token, jwt_required

app = Flask(__name__)

# Configurations
app.config["MONGO_URI"] = "mongodb://localhost:27017/flightstatus"
app.config["JWT_SECRET_KEY"] = "super-secret"  # Change this in production

# Initialize MongoDB and JWT
mongo = PyMongo(app)
jwt = JWTManager(app)

# Mock user for demonstration purposes
USERS = {"admin": "password123"}

@app.route('/api/login', methods=['POST'])
def login():
    username = request.json.get("username")
    password = request.json.get("password")
    if USERS.get(username) == password:
        access_token = create_access_token(identity=username)
        return jsonify(token=access_token), 200
    return jsonify({"msg": "Bad credentials"}), 401

@app.route('/api/flights', methods=['GET'])
@jwt_required()
def get_flights():
    flights = list(mongo.db.flights.find({}, {"_id": 0}))
    return jsonify(flights)

if __name__ == '__main__':
    app.run(debug=True)
