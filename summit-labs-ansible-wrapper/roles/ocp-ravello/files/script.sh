#!/bin/bash

ansible-playbook --inventory-file=/usr/share/ansible/summit/inventory/ /usr/share/ansible/summit/playbooks/ocp-pre-ravello.yml
ansible-playbook --inventory-file=/usr/share/ansible/summit/inventory/ /usr/share/ansible/summit/roles/applier/playbooks/openshift-cluster-seed.yml