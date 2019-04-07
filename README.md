# septa-slackbot
Utilize the septa api (http://www3.septa.org/hackathon/) to get updates on SEPTA

Currently implemented commands:
- `train status (train no)` get the next stop for a train and how late it is
- `get trains next to arrive at (station)` get trains next to arrive at a station
- `get all trains` get all train numbers and their respective line

======================================================================================

## Local Installation

### Prerequisites:
- [go](https://github.com/golang/go)
- [dep](https://github.com/golang/dep)

### Setup local.env
1. `$ cp local_env_template local.env`
2. Fill in local.env settings

### Start the project
1. `$ make setup`
2. `$ make compile`
3. `$ make run`

======================================================================================

## Running a Local Docker Image

### Prerequisites:
- [docker](https://docs.docker.com/install/)

1. `$ cp local_env_template local.env`
2. Fill in local.env settings
3. `$ make build-docker`
4. `$ make run-docker`

======================================================================================

## Running in Kubernetes

### Prerequisites:
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

1. `$ cp k8s/secret_template.yaml k8s/secret.yaml`
2. `$ echo -n 'slack-token' | base64`
3. Put resulting value in k8s/secret.yaml
4. `$ kubectl create -f k8s/secret.yaml`
5. `$ kubectl create -f k8s/deployment.yaml`

