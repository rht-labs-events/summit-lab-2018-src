## Source files for 'Effectively monitor and troubleshoot your OCP Cluster' Red Hat Summit 2018 Lab

### Structure

  - `ansible-docker:` ansible role to customize Docker storage on Cluster Nodes
  - `ansible-rh-subscription-manager:` ansoble role to customize Red Hat Subscriptions on Cluster Nodes
  - `ansible-ssh-setup:` ansible role to setup ssh keys and ssh trust from the bastion server onto all other servers
  - `ci-cd:` contains the repositories to bootstrap the ci-cd Pipelines that will trigger builds and deployments during the lab for "noise maker"
  - `lab-app:` contains the jenkins DSL definition to populate all the Jenkins jobs around ci-cd for the React app
  - `lab-react-app:` forked from `https://github.com/facebook/create-react-app`, a sample React FE app
  - `openshift-ansible:` custom openshift deployment playbooks with some modifications on the Prometheus and Grafana components
