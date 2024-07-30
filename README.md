# Flight Status and Notifications System

## Description
This project is a comprehensive Flight Status and Notifications System that provides real-time updates and notifications to passengers. The system includes features for displaying current flight status, sending notifications for flight status changes, and integrating with airport systems to pull accurate data. The frontend is designed with a modern, user-friendly interface inspired by the Indigo Airlines website.

## Features
- **Real-time Updates**: Displays current flight status including delays, cancellations, and gate changes.
- **Push Notifications**: Sends notifications for flight status changes via SMS, email, or app notifications using technologies like Kafka and RabbitMQ.
- **Integration with Airport Systems**: Pulls data from airport databases for accurate flight information.

## Technologies
- **Frontend**: HTML, CSS, React.js
- **Backend**: Python, Go, Java
- **Database**: MongoDB, PostgreSQL
- **Notifications**: Firebase Cloud Messaging, Kafka, RabbitMQ, etc.


## Setup and Installation

### Frontend
1. **Navigate to the frontend directory**:
    ```bash
    cd /ATCManagement/frontend
    ```
2. **Install dependencies**:
    ```bash
    npm install
    ```
3. **Run the development server**:
    ```bash
    npm start
    ```

### Backend (Python)
1. **Navigate to the Python backend directory**:
    ```bash
    cd /ATCManagement/backend/python
    ```
2. **Create a virtual environment**:
    ```bash
    python -m venv venv
    ```
3. **Activate the virtual environment**:
    - On Windows:
      ```bash
      venv\Scripts\activate
      ```
    - On MacOS/Linux:
      ```bash
      source venv/bin/activate
      ```
4. **Install dependencies**:
    ```bash
    pip install -r requirements.txt
    ```
5. **Run the application**:
    ```bash
    python app.py
    ```

### Backend (Go)
1. **Navigate to the Go backend directory**:
    ```bash
    cd /ATCManagement/backend/go
    ```
2. **Run the application**:
    ```bash
    go run main.go
    ```

### Backend (Java)
1. **Navigate to the Java backend directory**:
    ```bash
    cd /ATCManagement/backend/java
    ```
2. **Compile and run the application**:
    ```bash
    javac Main.java
    java Main
    ```

### Database
1. **MongoDB**:
    - Ensure MongoDB is installed and running.
    - Import the initial data:
      ```bash
      mongo < init-mongo.js
      ```
2. **PostgreSQL**:
    - Ensure PostgreSQL is installed and running.
    - Import the initial data:
      ```bash
      psql -U <username> -d <database> -f init-postgres.sql
      ```

## Usage
1. **Login**: Start by logging in with your username and password.
2. **Flight Status**: View the current status of flights, including any delays, cancellations, or gate changes.
3. **Notifications**: Receive real-time notifications about flight status changes.

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License
This project is licensed under the MIT License.
