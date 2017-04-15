pipeline {
  agent {
    docker {
      image 'golang'
    }
    
  }
  stages {
    stage('Build') {
      steps {
        parallel(
          "Build": {
            sh 'echo Hello'
            archiveArtifacts(artifacts: '*.tar.gz', onlyIfSuccessful: true)
            
          },
          "Test": {
            sh 'go test'
            
          }
        )
      }
    }
  }
}