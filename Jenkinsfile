pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS=credentials('docker-password')
    }
    stages {
        stage('build docker image') {
            steps {
                sh 'docker build -t webserver .'
            }
        }
        stage('create docker tag') {
            steps {
                sh 'docker tag webserver naczaaja/todoapp-final_webserver:latest'
            }
        }
        stage('docker login') {
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            }
        }
        stage('push to docker hub') {
            steps {
                sh 'docker push naczaaja/todoapp-final_webserver:latest'
            }
        }
        stage('test login to production server'){
            steps{
                sshagent(credentials:['private_key_jenkins']){
                    sh 'ssh  -o StrictHostKeyChecking=no root@164.92.84.214 '
                }
            }
        }
        stage('tranfer docker-compose to production'){
            steps{
                sshagent(credentials:['private_key_jenkins']){
                    sh 'scp -o StrictHostKeyChecking=no docker-compose.yml root@164.92.84.214:/root/docker-compose.yml'
                }
            }
        }
        stage('rebuild image an compose run'){
            steps{
                sshagent(credentials:['private_key_jenkins']){
                    sh 'ssh  -o StrictHostKeyChecking=no root@164.92.84.214 docker rmi naczaaja/todoapp-final_webserver -f'
                    sh 'ssh  -o StrictHostKeyChecking=no root@164.92.84.214 docker-compose up -d'
                }
            }
        }
    }
}