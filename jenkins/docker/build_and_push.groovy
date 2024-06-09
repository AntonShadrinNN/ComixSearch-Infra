node {
    stage('Preparation') {
        git(
            credentialsId: 'sys_ci',
            url: 'git@github.com:AntonShadrinNN/ComixSearch-Infra.git',
            branch: branch
        )

    }

    stage('Build') {
        sh("podman build -t ${image_name}/${image_tag} -f ${file_path}")
    }

    stage('Push') {
        if (push_to_repo.toBoolean()) {
            withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: account,
                               usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
                sh("podman login -u $USERNAME --password $PASSWORD ${registry}")
                sh("podman push ${image_name}:${image_tag} $USERNAME/${remote_repo}:${image_tag}")   
           }
        }
    }

}
