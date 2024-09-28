// Package env provides a constant environment variable for application.
//
// Note: For Kubernetes to create secrets, use the bash script "create_k8s_secret.sh" (located in the "k8s-deployment" directory).
// For example, put "DB_DATABASE=databasename", "DB_USERNAME=dbusername", "DB_PASSWORD=yourdbpasswordpogger" in a file named ".env" or whatever any other desired name.
//
// Example ".env" file:
//
//	DB_DATABASE=databasename
//	DB_USERNAME=dbusername
//	DB_PASSWORD=yourdbpasswordpogger
//
// Then run the "create_k8s_secret.sh" script to create the Kubernetes secrets.
// Note: If your Kubernetes cluster already has a built-in Hardware Security Module (HSM), you don't need to use an external secrets mechanism.
package env
