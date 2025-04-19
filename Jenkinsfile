pipeline {
    agent any
    stages {
        stage('Clonar repositorio') {
            steps {
                git 'https://github.com/EstebanCP2003/Registro-Veterinario.git'
            }
        }
        stage('Ejecutar pruebas') {
            steps {
                sh 'go test ./... -v -json > report.json'
            }
        }
        stage('Build') {
            steps {
                sh 'go build'
            }
        }
        post {
            always {
                junit 'report.json' // si generas reporte JUnit
            }
        }
    }
}