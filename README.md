# Sorting-Server
How to Run the project
1. Clone the project
```bash
    git clone https://github.com/suhail34/Sorting-Server.git
```
2. Change the directory
```bash
    cd Sorting-Server
```
3. For Running the Server locally
```bash
    make server-local
```
4. For Running the Server in a docker container
```bash
    make server-container-start
```
5. For Stopping the running docker container
```bash
    make server-container-stop
```
6. Check the endpoints using curl
For Sequential Processing
```bash
    curl -X POST -H "Content-Type:application/json" --data @body.json http://localhost:8000/process-single
```
For Concurrent Processing
```bash
    curl -X POST -H "Content-Type:application/json" --data @body.json http://localhost:8000/process-concurrent
```

