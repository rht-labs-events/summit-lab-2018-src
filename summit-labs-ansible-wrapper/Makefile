MAKEFLAGS += --silent

setup: setup-ansible

setup-ansible:
	ansible-galaxy install -r roles/requirements.yml -f -c

ansible-install:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-install.yml

ansible-pre:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-pre.yml

ansible-pre-ravello:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-pre-ravello.yml

ansible-post:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-post.yml

ansible-uninstall:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-uninstall.yml

ansible-install-glusterfs:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-install-glusterfs.yml

ansible-install-hosted:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-install-hosted.yml

ansible-install-ci-cd-lab:
	ansible-playbook --inventory-file=./inventory/ roles/applier/playbooks/openshift-cluster-seed.yml

ansible-install-prometheus:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-install-prometheus.yml

ansible-uninstall-prometheus:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-uninstall-prometheus.yml

ansible-install-grafana:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-install-grafana.yml

ansible-sync-time:
	ansible-playbook --inventory-file=./inventory/ playbooks/ocp-sync-time.yml

### Scenario 0

s0-start: s0-stop-etcd3 s0-stop-etcd2

s0-stop-etcd3:
	ansible-playbook --inventory-file=./inventory/ playbooks/s0-stop-etcd3.yml

s0-stop-etcd2:
	ansible-playbook --inventory-file=./inventory/ playbooks/s0-stop-etcd2.yml

s0-start-etcd3:
	ansible-playbook --inventory-file=./inventory/ playbooks/s0-start-etcd3.yml

s0-start-etcd2:
	ansible-playbook --inventory-file=./inventory/ playbooks/s0-start-etcd2.yml

s0-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s0-solve.yml

### Scenario 1

s1-start:
	ansible-playbook --inventory-file=./inventory/ playbooks/s1-start.yml

s1-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s1-solve.yml

### Scenario 2

s2-start:
	ansible-playbook --inventory-file=./inventory/ playbooks/s2-start.yml

s2-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s2-solve.yml

### Scenario 3

s3-start:
	ansible-playbook --inventory-file=./inventory/ playbooks/s3-start.yml

s3-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s3-solve.yml

### Scenario 4

s4-start:
	ansible-playbook --inventory-file=./inventory/ playbooks/s4-start.yml

s4-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s4-solve.yml

### Scenario 5

s5-start:
	ansible-playbook --inventory-file=./inventory/ playbooks/s5-stop-ovs.yml

s5-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s5-solve.yml

### Scenario 6

s6-start:
	ansible-playbook --inventory-file=./inventory/ playbooks/s6-stop-dnsmasq.yml

s6-solve:
	ansible-playbook --inventory-file=./inventory/ playbooks/s6-solve.yml
