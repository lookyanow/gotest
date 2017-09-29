pipeline {
  agent any
  stages {
    stage('goserver build') {
      steps {
        sh '''
        ./server-make.sh
'''
      }
    }
    stage('goserver container build') {
      steps {
        sh '''
        docker build -t goserver ./server/Dockerfile  
'''
   }
  }
 }
}
