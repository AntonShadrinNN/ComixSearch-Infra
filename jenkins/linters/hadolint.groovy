node {
    stage('Preparation') {
        git(
            credentialsId: 'sys_ci',
            url: 'git@github.com:AntonShadrinNN/ComixSearch-Infra.git',
            branch: branch
        )

    }

    String dockerfiles = ''
    stage('Find all dockerfiles') {
        dockerfiles = findFiles(glob: '**/Dockerfile*').collect{ file ->
            return file.path
        }.join(' ')
    }

    stage('Lint') {
        sh("hadolint -f json --no-color ${dockerfiles} > res.json")
    }

    stage('Archive output') {
        archiveArtifacts(
            artifacts: 'res.json'
        )
    }
}
