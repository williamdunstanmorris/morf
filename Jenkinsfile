// node("docker") {
//     docker.withRegistry('https://hub.docker.com/r/williamdunstanmorris/morf', 'williamdunstanmorris') {
//
//         git url: "https://github.com/williamdunstanmorris/morf", credentialsId: 'williamdunstanmorris'
//
//         sh "git rev-parse HEAD > .git/commit-id"
//         def commit_id = readFile('.git/commit-id').trim()
//         println commit_id
//
//         stage "build"
//         def app = docker.build "morf"
//
//         stage "publish"
//         app.push 'master'
//         app.push "${commit_id}"
//     }
// }
//

pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                echo 'I am running on EC2 Baby!'
            }
        }
        stage('test') {
            steps {
                sh 'go version'
            }
        }
        stage('package') {
            steps {
                sh 'go version'
            }
        }
        stage('install') {
            steps {
                sh 'go version'
            }
        }
    }
}
