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
        docker build -t goserver ./  
'''
   }
  }
 }
}
