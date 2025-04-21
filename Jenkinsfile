pipeline {
    agent {
        label 'go-agent' // o usa docker { image 'golang:1.21' }
    }
    stages {
        stage('Clonar repositorio') {
            steps {
                git 'https://github.com/EstebanCP2003/Registro-Veterinario.git'
            }
        }
        stage('Ejecutar pruebas') {
            steps {
                sh 'go test -v ./... | tee result.out | go-junit-report > report.xml'
            }
        }
        stage('Build') {
            steps {
                sh 'go build'
            }
        }
    }
    post {
        always {
            junit 'report.xml'
        }
    }
}
