pipeline {
    agent {
        label 'agent2' // Esto asegura que todo corre en agent2
    }
    stages {

        stage('Clonar repositorio') {
            steps {
                git branch: 'main', url: 'https://github.com/EstebanCP2003/Registro-Veterinario.git'
            }
        }

        stage('Ejecutar pruebas') {
            steps {
                sh '''
                    go install github.com/jstemmer/go-junit-report@latest
                    go test -v ./... | tee result.out
                    $HOME/go/bin/go-junit-report < result.out > report.xml
                '''
            }
        }
        stage('Compilar proyecto') {
            steps {
                sh 'go build'
            }
        }
    }
    post {
        always {
            junit 'report.xml' // Publica los resultados de pruebas
        }
    }
}
