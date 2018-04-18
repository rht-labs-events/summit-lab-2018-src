ocp-pre-hooks
=========

Role to execute actions on all openshift nodes before execution. In example do some tactical fixes, or change configuration.

Requirements
------------

Running openshift cluster

Role Variables
--------------

rhel7cis_host_allow:
  - "10.252.4.0/255.255.252.0"  
  - "10.217.0.1/255.255.0.0"  
  - "10.218.0.1/255.255.0.0"
  - "10.219.0.1/255.255.0.0"
  - "10.254.0.0/255.255.0.0"
  - "10.252.0.0/255.255.0.0"

Author Information
------------------

Mangirdas Judeikis m.j@redhat.com
