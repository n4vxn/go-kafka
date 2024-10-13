# Simple Kafka Notification Producer and Consumer

This project demonstrates a simple notification system implemented in Go using Apache Kafka for messaging. The application consists of a producer that generates random notification messages and sends them to a Kafka topic, and a consumer that reads these messages from the topic.

## Technologies Used

- Go: The programming language used for building the application.
- Kafka: A distributed messaging system that handles the message broker functionality.
- Docker: Used to containerize the application and the Kafka environment, making it easy to set up and run.
- Logrus: To log produced and consumed notification.

## Getting Started

To run this application, you need to have Docker and Docker Compose installed on your machine.

### Prerequisites

- Go 1.18 or later
- Docker
- Docker Compose

### Running the Application

1. Clone the repository:

   ```bash
   git clone https://github.com/n4vxn/go-kafka.git
   cd go-kafka

2. Start the Kafka and Zookeeper Services:

    docker-compose up -d

3. Run the Go application:
    go run main.go