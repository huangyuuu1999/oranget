pipeline {
    agent {
        label "local"  // 拉取、镜像制备、部署在同一节点
    } // 只在local 节点构建，部署

    stages {
        stage('build-image') {
        	echo "go build..."
		}
        stage('push-image&&pull') {
	        echo "pulling..."
        }
        stage('deploy') {
          	echo "deploy! running...."
        }
    }
}