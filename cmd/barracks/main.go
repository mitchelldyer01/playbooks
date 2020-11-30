package barracks

import {
	"fmt"
	deploy "internal/applications"
}



func Main() {
	applications := map[string]deploy.ImageConfig {
		"registry" : deploy.ImageConfig {
			tag				: "registry:2.7.1",
			hostIP			: "0.0.0.0",
			hostPort		: "5000",
			containerPort	: "5000",
			protocol		: "tcp",
		},
	}

	for _, application := range applications {
		deploy.Deploy(application)
	}
}
