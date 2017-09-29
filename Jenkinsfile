pipeline {
  agent any
  stages {
    stage('test') {
      steps {
        sh '''cd server/
go build server2.go 
'''
      }
    }
  }
}