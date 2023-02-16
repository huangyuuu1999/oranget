pipeline {
    agent {
        label "local"  // 拉取、镜像制备、部署在同一节点
    } // 只在local 节点构建，部署

    stages {
        stage('build-image') {
			steps {
                // 停止之前启动的容器
				script {
                    try {
                        sh """ 
                            docker stop app_back || true
                            docker rm app_back || true
                        """
                    } catch(Exception err) {
                        echo "no app running."
                    }
                }
                // 删除旧的镜像
                script {
                    try {
                    sh """ 
                        docker rmi 43.139.176.247/fruit_buckets/oranget:latest
                    """
                    } catch(Exception err) {
                        echo "no 43.139.176.247/fruit_buckets/oranget:latest image!"
                    }
                }
			}
        }
        stage('build && push-image && pull') {
			steps{
	        	echo 'Building..已经拉取源码，在此处执行构建'// 已经拉取源码，在此处执行构建
                withCredentials([usernamePassword(credentialsId: 'jenkins-access-harbor', passwordVariable: 'PASSWORD', usernameVariable: 'USERNAME')]){
                    sh """
                        echo ${PASSWORD} | docker login -u 'gongyulei' --password-stdin 43.139.176.247/fruit_buckets

                        # build
                        docker build -t 43.139.176.247/fruit_buckets/oranget:latest .
                        docker push 43.139.176.247/fruit_buckets/oranget:latest
                    """
                }
			}
        }
        stage('build cronjob image') {
            when {
                anyOf {
                    changeset 'src/scrapy/*'
                }
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'jenkins-access-harbor', passwordVariable: 'PASSWORD', usernameVariable: 'USERNAME')]){
                    sh """
                        echo ${PASSWORD} | docker login -u 'gongyulei' --password-stdin 43.139.176.247/fruit_buckets

                        # build
                        docker build -t 43.139.176.247/fruit_buckets/oranjob:latest ./src/scrapy/Dockerfile
                        docker push 43.139.176.247/fruit_buckets/oranjob:latest
                    """
                }
            }
        }
        stage('deploy') {
			steps{
          		withCredentials([usernamePassword(credentialsId: 'jenkins-access-harbor', passwordVariable: 'PASSWORD', usernameVariable: 'USERNAME')]){
                    sh """
                        # echo ${PASSWORD} | docker login -u 'gongyulei' --password-stdin 43.139.176.247/fruit_buckets

                        # pull
                        # docker pull 43.139.176.247/fruit_buckets/oranget:latest
                        # deploy
                        # docker run -d --name app_back -p4000:8080 43.139.176.247/fruit_buckets/oranget:latest
                        echo 'do not apply by docker anymore...'
                    """
                }
	    }
        }
    }
}
