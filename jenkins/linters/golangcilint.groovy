node {
    stage('Preparation') {
        git(
            credentialsId: 'sys_ci',
            url: 'git@github.com:AntonShadrinNN/ComixSearch-Infra.git',
            branch: branch
        )

    }

    Integer res = 0
    stage('Lint') {
        dir(service_path) {
            try {
                sh('golangci-lint run > res.txt')
            } catch(err) {
                res = 1
            }
        }
    }

    stage('Archive output') {
        archiveArtifacts(
            artifacts: '**/res.txt'
        )
    }

    return res
}
