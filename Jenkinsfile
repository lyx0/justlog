pipeline {
  agent any
  stages {
    stage('TestBuild') {
      steps {
        sh 'go build'
      }
    }
    stage('Testing') {
      steps {
        sh 'go test'
      }
    }
    stage('Docker Build') {
      steps {
        sh '''docker build .
docker tag gempbotgo_gempbotgo gempir/gempbotgo
docker push gempir/gempbotgo
	'''
      }
    }
    stage('Prepare Env') {
      steps {
        sh 'cp ./prod.yml /home/gempir/gempbotgo'
      }
    }
    stage('Deploy') {
      steps {
        sh '''cd /home/gempir/gempbotgo
docker-compose pull
docker-compose up'''
      }
    }
  }
}