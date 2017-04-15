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
            sh 'go get'
            archiveArtifacts(artifacts: '$GOPATH/bin/hd', onlyIfSuccessful: true)
          }
        )
      }
    }
  }
}
