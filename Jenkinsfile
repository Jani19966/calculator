pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh '''
                    # Install Golang
                    apt-get update && apt-get install -y golang

                    # Set up GOPATH and PATH
                    export GOPATH=$WORKSPACE/go
                    export PATH=$PATH:$GOPATH/bin

                    # Build the application
                    go build main.go
                '''
            }
        }
    }
}
