ansible-rh-subscription-manager
=========

Role subscribes to Red Hat subscription manager, attached pool and enables requires repositories

Role Variables
--------------

```
rh_username - RH username
rh_password - Password
rh_pool - pool it to be attached
rh_repos - repos list to be enabled
```

ToDo:
-----

  * Add pool listing using search by name
  * Add ability to use different repos and pools on differnet servers
  * Add motd to be optional
  * Add vault support for password

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for users too:

    - hosts: servers
      roles:
         - ansible-rh-subscription-manager


Author Information
------------------

Mangirdas[et]Judeikis[dot]LT