pipeline {
  agent any
  stages {
    stage('goserver build') {
      steps {
        sh '''
        go --version
        echo "Hello world from jenkins project"
'''
      }
    }
    stage('testing stage') {
      steps {
        sh '''
        echo "This is test stage" 
'''
   }
  }
 }
}
