pipeline {
  agent any
  stages {
    stage('test') {
      steps {
        sh '''
        ./server-make.sh
'''
      }
    }
  }
}
