node {
    def root = tool(
        type: 'go',
        name: '1.22.0'
    )

    stage('Preparation') {
        git(
            credentialsId: 'sys_ci',
            url: 'git@github.com:AntonShadrinNN/ComixSearch-Application.git',
            branch: branch
        )

    }

    stage('Build') {
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            if (install_deps) {
                sh("go mod tidy")
            }
            sh("go build -o main ${main_path}/${main_name}")

        }
    }

    stage('Test') {
        sh('go test -coverprofile cover.out ./...')
        sh('go tool cover -html=cover.out -o cov.html')
    }

    stage('Archive artifacts') {
        archiveArtifacts(
            artifacts: 'main, cov.html',
            allowEmptyArchive: false,
            fingerprint: false,
            onlyIfSuccessful: true
        )
    }

}
