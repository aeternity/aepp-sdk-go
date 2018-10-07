pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile.ci'
      args '-v /etc/group:/etc/group:ro ' +
           '-v /etc/passwd:/etc/passwd:ro ' +
           '-v /var/lib/jenkins:/var/lib/jenkins ' +
           '-v /usr/bin/docker:/usr/bin/docker:ro ' +
           '--network=host'
    }
  }

  environment {
    DOCKER_COMPOSE = "docker-compose -p ${env.BUILD_TAG} -H 127.0.0.1:2376"
  }

  stages {
    stage('Pre Test') {
      steps {
        sh 'go version'
        sh 'go get github.com/tebeka/go2xunit'
        sh 'go get -u github.com/golang/lint/golint'
        sh 'go mod download'
      }
    }

    stage('Test') {
      steps {
          withCredentials([usernamePassword(credentialsId: 'genesis-wallet',
                                          usernameVariable: 'WALLET_PUB',
                                          passwordVariable: 'WALLET_PRIV')]) {
          sh "${env.DOCKER_COMPOSE} run sdk sh -c 'make lint && make test' "
        }
      }
    }

    stage('Build') {
      steps {
          sh "${env.DOCKER_COMPOSE} run sdk make build-dist"
        }
      }
    }
  }

  post {
    always {
      junit 'test-results.xml'
      archive 'dist/*'
      sh 'docker-compose -H localhost:2376 down -v ||:'
    }
  }
}