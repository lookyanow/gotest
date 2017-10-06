pipeline {
  agent any
  stages {
    stage('goserver build') {
      steps {
        sh '''
        server-make.sh
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
