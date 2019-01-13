#!groovy

stage('Pre Test'){
    echo 'Before Test'
    sh 'go version'
    sh 'go get -u'
}

stage('Build'){
    echo 'Building Executable'
    sh """cd $GOPATH/src/github.com/DancesportSoftware/das"""
    sh """go build"""
}