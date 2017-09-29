pipeline {
  agent any
  stages {
    stage('goserver build') {
      steps {
        sh '''
        go build server.go
'''
      }
    }
    stage('goserver container build') {
      steps {
        sh '''
        sudo docker build -t goserver ./  
'''
   }
  }
 }
}
