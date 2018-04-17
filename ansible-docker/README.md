Docker Storage Configuration Ansible Role
=========

Role sole purpose is to configure docker to be used in multi-tenancy based environments. It will configure bigger pool space and storage configuration for `/var/lib/docker`

Requirements
------------

Additional disk need to be allocated to the server for this role to work. 

Role Variables
--------------

A description of the settable variables for this role should go here, including any variables that are in defaults/main.yml, vars/main.yml, and any variables that can/should be set via parameters to the role. Any variables that are read from other roles and/or the global scope (ie. hostvars, group vars, etc.) should be mentioned here as well.

```
docker_data : boolean variable is used to define if disk provisioned needs a lvm for /opt/docker/data. By Default this is set to false, so that there is no impact on Open Shift.
set this to true from playbook or extra vars for  /opt/docker/data lvm 
docker_storage_disk: /dev/sdb - addional disk for docker to use
docker_vg_name: docker_vg  - name of VG for docker
docker_lv1_name: usrdatadocker_lv - name of LV for data docker dir /opt/docker/data (hostPath or other cases)
docker_lv1_size: +50%FREE - How much space it should take from all disk.
docker_lv2_name: varlibdocker_lv - name of LV for /var/lib/docker
docker_lv2_size: +50%FREE  - How much space it should take from all disk.
docker_pool_size: +90%FREE - Size of docker_pool 
docker_data: false - Do we need to create LV1 (data dir)
```

Example, if you allocate 100GB aditional disk to docker and set values:
```
docker_lv1_size: +50%FREE
docker_lv2_size: +50%FREE
docker_pool_size: +90%FREE 
docker_data: true
```

result will be:
```
/opt/docker/data - 50GB
/var/lib/docker - 25GB
docker-pool - 22.5GB
```

Example Playbook
----------------

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for users too:

    - hosts: servers
      roles:
         - docker

Author Information
------------------

mangirdas[et]judeikis[dot]LT