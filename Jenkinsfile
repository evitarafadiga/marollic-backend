pipeline {
    agent any

    environment {
        DOCKER_COMPOSE_VERSION = '1.29.2'
        DOCKER_TLS_VERIFY = '0'
        DOCKER_HOST = 'tcp://docker:2375'
    }

    stages {
        state('verify') {
            steps {
                sh '''
                    docker version
                    docker info
                    docker compose version
                    curl --version
                    jq --version
                '''
            }
        }
        stage('Build') {
            steps {
                script {
                    sh 'docker build -f Dockerfile -t go-api-marollic .'
                }
            }
        }

        stage('Run') {
            steps {
                script {
                    sh 'docker run -d --name go-api-marollic go-api-marollic'
                }
            }
        }
    }

    stages {
        stage('Setup Docker Compose') {
            steps {
                script {
                    sh '''
                        docker --version
                        docker-compose --version
                    '''
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    sh 'go clean'
                    sh 'go mod tidy'
                    sh 'go build -o go-api/cmd/main.out go-api/cmd'
                }
            }
        }

        stage('Docker Compose Up') {
            steps {
                script {
                    sh 'docker-compose -f docker-compose.yml up -d'
                    sh 'docker-compose ps'
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    sh 'go test ./...'
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    sh 'docker-compose down'
                }
            }
        }
    }
}
